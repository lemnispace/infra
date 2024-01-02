package main

import (
	"testing"
)

const SECRET = "It's a Secret to Everybody"
const PAYLOAD = "Hello, World!"

func TestVerifySignature(t *testing.T) {
	event := LambdaEvent{
		Body: []byte(PAYLOAD),
		Headers: Headers{
			// https://docs.github.com/en/webhooks/using-webhooks/validating-webhook-deliveries#testing-the-webhook-payload-validation
			XHubSignature256: "sha256=757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
		},
	}
	result, err := verifySignature(event, SECRET)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !result {
		t.Error("signature verification failed")
	}
}
