package trigger

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
	s "github.com/lemnispace/infra/cmd/deploy/secret"
)

/*
Function to get the installation ID using the JWT

The installation ID is required to get the installation access token
*/
func GetInstallationID(ctx context.Context, client *github.Client) (int64, error) {
	// TODO: get org name from env
	resp, _, err := client.Organizations.ListInstallations(ctx, "lemnispace", nil)
	if err != nil {
		return 0, err
	}
	return resp.Installations[0].GetID(), nil
}

/*
Function to get the installation access token using the JWT

learn more at https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/authenticating-as-a-github-app-installation#generating-an-installation-access-token
*/
func getInstallationAccessToken(ctx context.Context, jwt string) (string, error) {
	client := github.NewClient(nil).WithAuthToken(jwt)
	id, err := GetInstallationID(ctx, client)
	if err != nil {
		return "", fmt.Errorf("error getting installation ID: %v", err)
	}
	token, _, err := client.Apps.CreateInstallationToken(ctx, id, nil)
	if err != nil {
		return "", err
	}
	return *token.Token, nil
}

func genJWT(appId string, pk *rsa.PrivateKey) (string, error) {
	// https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-json-web-token-jwt-for-a-github-app#about-json-web-tokens-jwts
	payload := jwt.MapClaims{
		"iat": time.Now().Add(time.Second * -60).Unix(), // subtract 60 seconds to account for clock skew
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

func getPk(ctx context.Context, ssm *s.SSM) (*rsa.PrivateKey, error) {
	// TODO: get param name from env
	pkPem, err := ssm.Param("/Any/infra/github-lemnispace-app-private-key", true).GetValue(ctx)
	if err != nil {
		return nil, err
	}
	return decodePk(pkPem)

}

func getAppId(ctx context.Context, ssm *s.SSM) (string, error) {
	// TODO: get param name from env
	appId, err := ssm.Param("/Any/infra/github-lemnispace-app-id", true).GetValue(ctx)
	if err != nil {
		return "", err
	}
	return appId, nil
}

func getAuthToken(ctx context.Context, ssm *s.SSM) (string, error) {
	pk, err := getPk(ctx, ssm)
	if err != nil {
		return "", fmt.Errorf("unable to get github app private key: %v", err)
	}
	appId, err := getAppId(ctx, ssm)
	if err != nil {
		return "", fmt.Errorf("unable to get github app id: %v", err)
	}
	token, err := genJWT(appId, pk)
	if err != nil {
		return "", fmt.Errorf("unable to generate jwt: %v", err)
	}
	return getInstallationAccessToken(ctx, token)
}

func getResponse(body io.ReadCloser) (string, error) {
	b, err := io.ReadAll(body)
	defer body.Close()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func TriggerDeploy(ctx context.Context, ssm *s.SSM) error {
	token, err := getAuthToken(ctx, ssm)
	if err != nil {
		return fmt.Errorf("unable to get github app token: %v", err)
	}
	owner := os.Getenv("DEPLOYMENT_REPO_OWNER")
	repo := os.Getenv("DEPLOYMENT_REPO_NAME")
	deployFile := os.Getenv("DEPLOYMENT_FILE_NAME")
	client := github.NewClient(nil).WithAuthToken(token)
	resp, err := client.Actions.CreateWorkflowDispatchEventByFileName(ctx, owner, repo, deployFile, github.CreateWorkflowDispatchEventRequest{Ref: "main"})
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
