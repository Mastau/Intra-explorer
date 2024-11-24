package api42

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const (
	authURL    = "https://api.intra.42.fr/oauth/token"
)

func init_api42() (string, string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", "", fmt.Errorf("Error loading .env file: %v", err)
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		return "", "", fmt.Errorf("Error set client id and client secret in the .env file")
	}
	return clientID, clientSecret, nil
}

func GetAccessToken() (string, error) {

	clientID, clientSecret, err := init_api42()
	if err != nil {
		return "", err
	}

	client := resty.New() // create new http client instance 

	resp, err := client.R(). // init new request
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     clientID,
			"client_secret": clientSecret,
		}).
		Post(authURL)

	if err != nil {
		return "", fmt.Errorf("error getting access token: %v", err)
	}

	var authResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &authResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling auth response: %v", err)
	}

	accessToken, ok := authResponse["access_token"].(string)
	if !ok || accessToken == "" {
		return "", fmt.Errorf("no access token found in the response")
	}

	return accessToken, nil
}
