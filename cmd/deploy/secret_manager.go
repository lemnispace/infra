package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSM struct {
	client *ssm.Client
}

func NewSSMClient(ctx context.Context) (*SSM, error) {
	// Load the AWS default config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}
	client := ssm.NewFromConfig(cfg)
	return &SSM{client}, nil
}

type Param struct {
	Name           string
	WithDecryption bool
	ssmsvc         *SSM
}

// Param creates the struct for querying the param store
func (s *SSM) Param(name string, decryption bool) *Param {
	return &Param{
		Name:           name,
		WithDecryption: decryption,
		ssmsvc:         s,
	}
}

func (p *Param) GetValue(ctx context.Context) (string, error) {
	input := &ssm.GetParameterInput{
		Name:           &p.Name,
		WithDecryption: &p.WithDecryption,
	}
	parameter, err := p.ssmsvc.client.GetParameter(ctx, input)
	if err != nil {
		return "", err
	}
	return *parameter.Parameter.Value, nil
}
