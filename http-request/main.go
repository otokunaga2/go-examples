package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var apiAccessToken string

//// ポイント1
type Sample struct {
	Sentence string `json:"sentence"`
	Type     string `json:"type"`
}
type APITokenRequest struct {
	GrantType    string `json:"grantType"`
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type APITokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
	IssuedAt    string `json:"issued_at"`
}

func getAPIToken() (string, error) {
	client := &http.Client{}
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	jsonData, err := json.Marshal(APITokenRequest{GrantType: "client_credentials", ClientId: clientID, ClientSecret: clientSecret})
	fmt.Println("Input data %s", string(jsonData))

	if err != nil {
		fmt.Errorf("Fail to marshal json %s", err)
	}
	req, _ := http.NewRequest(
		"POST",
		"https://api.ce-cotoha.com/v1/oauth/accesstokens",
		bytes.NewBuffer(jsonData),
	)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	var decodedJson APITokenResponse
	defer resp.Body.Close()
	fmt.Println("Get response from API is %s", string(body))
	if err := json.Unmarshal([]byte(body), &decodedJson); err != nil {
		fmt.Errorf("Fail to marshal json %s", err)
		return "", err
	}
	return decodedJson.AccessToken, nil
}

// Refresh if api seemed outdated
func refreshAPIKey() (string, error) {
	//var refreshed string
	if apiAccessToken == "" {
		fmt.Println("Pass here")
		refreshed, err := getAPIToken()
		if err != nil {
			//os.Exit(1)
			return apiAccessToken, err
		}
		apiAccessToken = refreshed
	}
	return apiAccessToken, nil

}

func main() {
	var apiToken string
	if result, err := refreshAPIKey(); err != nil {
		fmt.Errorf("Fail to refreshAPI ")
		os.Exit(1)
	} else {
		//fmt.Println("Resutl %s", result)
		apiToken = result
	}

	client := &http.Client{}
	jsonData, err := json.Marshal(Sample{Type: "default", Sentence: "体温 36.5"})
	fmt.Println("Input data %s", string(jsonData))

	if err != nil {
		fmt.Errorf("Fail to marshal json %s", err)
	}
	req, _ := http.NewRequest(
		"POST",
		"https://api.ce-cotoha.com/api/dev/nlp/v1/sentence_type",
		bytes.NewBuffer(jsonData),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println("Get response from API is %s", string(body))
}
