package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	Body            string  `json:"body"`
	Headers         Headers `json:"headers"`
	IsBase64Encoded bool    `json:"isBase64Encoded"`
}

type Headers struct {
	Accept                    string `json:"accept"`
	AcceptEncoding            string `json:"accept-encoding"`
	ContentLength             string `json:"content-length"`
	ContentType               string `json:"content-type"`
	XGithubDelivery           string `json:"x-github-delivery"`
	XGithubEvent              string `json:"x-github-event"`
	XGithubHookID             string `json:"x-github-hook-id"`
	XGithubHookInstallationID string `json:"x-github-hook-installation-target-id"`
	XGithubHookType           string `json:"x-github-hook-installation-target-type"`
	XHubSignature             string `json:"x-hub-signature"`
	XHubSignature256          string `json:"x-hub-signature-256"`
}

type WebhookEvent struct {
	Zen        string     `json:"zen"`
	HookID     int        `json:"hook_id"`
	Hook       Hook       `json:"hook"`
	Repository Repository `json:"repository"`
}

type Repository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
}

type Hook struct {
	Type          string     `json:"type"`
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Active        bool       `json:"active"`
	Events        []string   `json:"events"`
	Config        HookConfig `json:"config"`
	UpdatedAt     string     `json:"updated_at"`
	CreatedAt     string     `json:"created_at"`
	URL           string     `json:"url"`
	TestURL       string     `json:"test_url"`
	PingURL       string     `json:"ping_url"`
	DeliveriesURL string     `json:"deliveries_url"`
}

type HookConfig struct {
	ContentType string `json:"content_type"`
	InsecureSSL string `json:"insecure_ssl"`
	Secret      string `json:"secret"`
	URL         string `json:"url"`
}

func getPayload(event LambdaEvent) ([]byte, error) {
	if event.IsBase64Encoded {
		var err error
		payload, err := base64.StdEncoding.DecodeString(event.Body)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("error decoding base64 payload: %s", err)
		}
		return payload, nil
	}
	return []byte(event.Body), nil
}

func verifySignature(sig string, payload []byte, secret string) (bool, error) {
	expectedMAC := calculateHash(secret, payload)
	receivedMAC, err := hex.DecodeString(strings.TrimPrefix(sig, "sha256="))
	if err != nil {
		log.Println(err)
		return false, fmt.Errorf("error decoding received signature: %s", err)
	}
	return hmac.Equal(expectedMAC, receivedMAC), nil
}

func calculateHash(secret string, payload []byte) []byte {
	signature := hmac.New(sha256.New, []byte(secret))
	signature.Write(payload)
	return signature.Sum(nil)
}

func getSSM(ctx context.Context) (*SSM, error) {
	ssm, err := NewSSMClient(ctx)
	if err != nil {
		return nil, err
	}
	return ssm, nil
}

func getSecret(ctx context.Context, ssm *SSM) (string, error) {
	secret, err := ssm.Param("/Any/infra/webhook-secret", true).GetValue(ctx)
	if err != nil {
		return "", err
	}
	return secret, nil
}

func decodePayload(bp []byte) (*WebhookEvent, error) {
	p := strings.TrimPrefix(string(bp), "payload=")
	decoded, err := url.QueryUnescape(p)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error decoding payload: %s", err)
	}
	var w WebhookEvent
	json.Unmarshal([]byte(decoded), &w)
	return &w, nil
}

func HandleRequest(ctx context.Context, event LambdaEvent) (string, error) {
	ssm, err := getSSM(ctx)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error getting SSM client")
	}
	secret, err := getSecret(ctx, ssm)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error getting secret")
	}
	payload, err := getPayload(event)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error processing payload")
	}
	valid, err := verifySignature(event.Headers.XHubSignature256, payload, secret)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error verifying signature")
	}
	if !valid {
		return "", fmt.Errorf("invalid signature")
	}
	wh, err := decodePayload(payload)
	if err != nil || wh == nil {
		log.Print(err)
		return "", fmt.Errorf("invalid payload")
	}
	if wh.Hook.Events[0] == "deployment" {
		log.Printf("Deployment webhook received from %s", wh.Repository.FullName)
		err := TriggerDeploy(ctx, ssm)
		if err != nil {
			log.Printf("Error in deployment: %v", err)
		}
	}
	return "webhook received", nil
}

func main() {
	lambda.Start(HandleRequest)
}
