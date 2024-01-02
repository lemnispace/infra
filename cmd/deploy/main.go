package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type WebhookEvent struct {
	Zen         string `json:"zen"`
	HookID      int    `json:"hook_id"`
	Hook        Hook   `json:"hook"`
}

type Hook struct {
	Type        string       `json:"type"`
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Active      bool         `json:"active"`
	Events      []string     `json:"events"`
	Config      HookConfig   `json:"config"`
	UpdatedAt   string       `json:"updated_at"`
	CreatedAt   string       `json:"created_at"`
	URL         string       `json:"url"`
	TestURL     string       `json:"test_url"`
	PingURL     string       `json:"ping_url"`
	DeliveriesURL string     `json:"deliveries_url"`
	LastResponse LastResponse `json:"last_response"`
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

func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	log.Print("Received context: ", ctx)
	log.Println("Received event: ", event)
	err_msg := "Error retrieving secret for webhook validation"
	ssmsvc, err := NewSSMClient(ctx)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf(err_msg)
 }
	_, err = ssmsvc.Param("/Any/infra/webhook-secret", true).GetValue(ctx)
	if err != nil {
		 log.Println(err)
		 return "", fmt.Errorf(err_msg)
	}
	return "success", nil
}

func main() {
	lambda.Start(HandleRequest)
}
