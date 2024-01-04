package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/go-github/v57/github"
)

func genJWT(appId string, pk *rsa.PrivateKey) (string, error) {
	// https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-json-web-token-jwt-for-a-github-app#about-json-web-tokens-jwts
	payload := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 10).Unix(),
		"iss": appId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)
	jwtString, err := token.SignedString(pk)
	if err != nil {
		return "", err
	}
	return jwtString, nil
}

func decodePk(pkPem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pkPem))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}
	pk, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pk, nil
}

func getPk(ctx context.Context, ssm *SSM) (*rsa.PrivateKey, error) {
	pkPem, err := ssm.Param("/Any/infra/github-lemnispace-app-private-key", true).GetValue(ctx)
	if err != nil {
		return nil, err
	}
	return decodePk(pkPem)

}

func getAppId(ctx context.Context, ssm *SSM) (string, error) {
	appId, err := ssm.Param("/Any/infra/github-lemnispace-app-id", true).GetValue(ctx)
	if err != nil {
		return "", err
	}
	return appId, nil
}

func getAuthToken(ctx context.Context, ssm *SSM) (string, error) {
	pk, err := getPk(ctx, ssm)
	if err != nil {
		return "", fmt.Errorf("unable to get github app private key: %v", err)
	}
	appId, err := getAppId(ctx, ssm)
	if err != nil {
		return "", fmt.Errorf("unable to get github app id: %v", err)
	}
	return genJWT(appId, pk)
}

func getResponse(body io.ReadCloser) (string, error) {
	b, err := io.ReadAll(body)
	defer body.Close()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func TriggerDeploy(ctx context.Context, ssm *SSM) error {
	token, err := getAuthToken(ctx, ssm)
	if err != nil {
		return fmt.Errorf("unable to get github app token: %v", err)
	}
	owner := os.Getenv("DEPLOYMENT_REPO_OWNER")
	repo := os.Getenv("DEPLOYMENT_REPO_NAME")
	deployFile := os.Getenv("DEPLOYMENT_FILE_NAME")
	client := github.NewClient(nil).WithAuthToken(token)
	resp, err := client.Actions.CreateWorkflowDispatchEventByFileName(ctx, owner, repo, deployFile, github.CreateWorkflowDispatchEventRequest{})
	if err != nil {
		return fmt.Errorf("unable to trigger deploy: %v", err)
	}
	body, err := getResponse(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body after triggering deploy: %v", err)
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("recieved %d error response from github: %s", resp.StatusCode, body)
	}
	log.Println("Successfully triggered deploy. Github response: ", body)
	return nil
}
