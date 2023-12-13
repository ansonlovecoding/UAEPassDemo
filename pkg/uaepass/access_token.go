package uaepass

import (
	http_client "UAEPassDemo/pkg/http"
	"encoding/json"
	"errors"
	"fmt"
)

type GetAccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	ExpireIn         int    `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// GetAccessToken
// url The url of access token api
// credential the credential for the specific clientID
// grantType Must have the value as “authorization_code”.
// redirectURI should be same value as you pass in the access code api
// code the code you get from access code api
func GetAccessToken(url, credential, grantType, redirectURI, code string) (string, error) {
	tokenURL := fmt.Sprintf(
		"%s?grant_type=%s&redirect_uri=%s&code=%s",
		url,
		grantType,
		redirectURI,
		code,
	)
	resp, err := http_client.Post(tokenURL, credential, nil, 60)
	if err != nil {
		return "", err
	}
	tokenResp := &GetAccessTokenResponse{}
	err = json.Unmarshal(resp, tokenResp)
	if err != nil {
		return "", err
	}

	if len(tokenResp.Error) > 0 {
		return "", errors.New(tokenResp.Error)
	}
	return tokenResp.AccessToken, nil
}
