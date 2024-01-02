package main

import (
	"encoding/base64"
	"testing"
)

// https://docs.github.com/en/webhooks/using-webhooks/validating-webhook-deliveries#testing-the-webhook-payload-validation
const SECRET = "It's a Secret to Everybody"
const PAYLOAD = "Hello, World!"
const SIG = "sha256=757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17"

func TestVerifySignature(t *testing.T) {
	result, err := verifySignature(SIG, []byte(PAYLOAD), SECRET)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !result {
		t.Error("signature verification failed")
	}
}

func TestGetPayload(t *testing.T) {
	event := LambdaEvent{
		Body: PAYLOAD,
		Headers: Headers{
			XHubSignature256: SIG,
		},
	}
	payload, err := getPayload(event)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(payload) != PAYLOAD {
		t.Errorf("unexpected payload: %s", payload)
	}
}

func TestGetPayloadBase64(t *testing.T) {
	event := LambdaEvent{
		Body: base64.StdEncoding.EncodeToString([]byte(PAYLOAD)),
		Headers: Headers{
			XHubSignature256: SIG,
		},
		IsBase64Encoded: true,
	}
	payload, err := getPayload(event)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if string(payload) != PAYLOAD {
		t.Errorf("unexpected payload: %s", payload)
	}
}
