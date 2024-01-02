package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	fmt.Println("Event:", event)
	fmt.Println("Context:", ctx)

	return "Hello World!", nil
}

func main() {
	lambda.Start(HandleRequest)
}
