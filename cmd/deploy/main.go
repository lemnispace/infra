package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

// LambdaEvent represents the structure of your Lambda event
type LambdaEvent struct {
	Body    []byte  `json:"body"`
	Headers Headers `json:"headers"`
}

// Headers represents the specific structure of your headers
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
	Zen    string `json:"zen"`
	HookID int    `json:"hook_id"`
	Hook   Hook   `json:"hook"`
}

type Hook struct {
	Type          string       `json:"type"`
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	Active        bool         `json:"active"`
	Events        []string     `json:"events"`
	Config        HookConfig   `json:"config"`
	UpdatedAt     string       `json:"updated_at"`
	CreatedAt     string       `json:"created_at"`
	URL           string       `json:"url"`
	TestURL       string       `json:"test_url"`
	PingURL       string       `json:"ping_url"`
	DeliveriesURL string       `json:"deliveries_url"`
	LastResponse  LastResponse `json:"last_response"`
}

type HookConfig struct {
	ContentType string `json:"content_type"`
	InsecureSSL string `json:"insecure_ssl"`
	Secret      string `json:"secret"`
	URL         string `json:"url"`
}

type LastResponse struct {
	Code    interface{} `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func verifySignature(event LambdaEvent, secret string) (bool, error) {
	expectedMAC := calculateHash(secret, event.Body)
	receivedMAC, err := hex.DecodeString(strings.TrimPrefix(event.Headers.XHubSignature256, "sha256="))
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

func HandleRequest(ctx context.Context, event LambdaEvent) (string, error) {
	errMsg := "Error retrieving secret for webhook validation"
	ssmsvc, err := NewSSMClient(ctx)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf(errMsg)
	}
	secret, err := ssmsvc.Param("/Any/infra/webhook-secret", true).GetValue(ctx)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf(errMsg)
	}
	valid, err := verifySignature(event, secret)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error verifying signature: %s", err)
	}
	if !valid {
		return "", fmt.Errorf("invalid signature")
	}
	var eventBody WebhookEvent
	json.Unmarshal(event.Body, &eventBody)
	log.Printf("Received Webhook Type: %s", eventBody.Hook.Type)
	return "success", nil
}

func main() {
	lambda.Start(HandleRequest)
}
