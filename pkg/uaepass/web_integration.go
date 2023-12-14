package uaepass

import (
	"UAEPassDemo/pkg/config"
	http_client "UAEPassDemo/pkg/http"
	"UAEPassDemo/pkg/redis"
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

type GetUserInfoResponse struct {
	Sub           string `json:"sub"`
	FullnameAR    string `json:"fullnameAR"`
	Gender        string `json:"gender"`
	Mobile        string `json:"mobile"`
	LastnameEN    string `json:"lastnameEN"`
	FullnameEN    string `json:"fullnameEN"`
	UUID          string `json:"uuid"` // UAE PASS Unique User Identifier
	LastnameAR    string `json:"lastnameAR"`
	IDN           string `json:"idn"` // Verified Emirates ID number, only Residents have idn, visitors doesn't have
	NationalityEN string `json:"nationalityEN"`
	FirstnameEN   string `json:"firstnameEN"`
	UserType      string `json:"userType"` // Level of assurance,ALL (SOP1, 2 & 3)
	NationalityAR string `json:"nationalityAR"`
	FirstnameAR   string `json:"firstnameAR"`
	Email         string `json:"email"`
	ProfileType   string `json:"profileType"` // 1 (Citizen or Residents), 2 (Visitors)
	UnifiedID     string `json:"unifiedID"`   // only for visitors,Unique id of the user and will be returned for all users
}

// Get the access code first for the next step
// https://docs.uaepass.ae/guides/authentication/web-application/1.-obtaining-the-oauth2-access-code
func RequestAccessCodeURL(env, responseType, redirectURI, state, scope, acrValues, language string) (string, error) {
	var url string
	var clientID string
	if env == "staging" {
		url = config.LocalConfig.Endpoints.Staging.Authorization
		clientID = config.LocalConfig.Endpoints.Staging.ClientID
	} else if env == "production" {
		url = config.LocalConfig.Endpoints.Production.Authorization
		clientID = config.LocalConfig.Endpoints.Production.ClientID
	} else {
		return "", errors.New("the parameter 'env' was invalid")
	}

	requestURL := fmt.Sprintf(
		"%s?response_type=%s&redirect_uri=%s&client_id=%s&state=%s&scope=%s&acr_values=%s&ui_locales=%s",
		url,
		responseType,
		redirectURI,
		clientID,
		state,
		scope,
		acrValues,
		language,
	)
	return requestURL, nil
}

// GetAccessToken
// url The url of access token api
// credential the credential for the specific clientID
// grantType Must have the value as “authorization_code”.
// redirectURI should be same value as you pass in the access code api
// code the code you get from access code api
func GetAccessToken(env, state, grantType, redirectURI, code string) (string, error) {
	var credential string
	var url string
	if env == "staging" {
		credential = config.LocalConfig.Endpoints.Staging.Credentials
		url = config.LocalConfig.Endpoints.Staging.Token
	} else if env == "production" {
		credential = config.LocalConfig.Endpoints.Production.Credentials
		url = config.LocalConfig.Endpoints.Production.Token
	} else {
		return "", errors.New("the parameter 'env' was invalid")
	}

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

	//store the token to redis
	err = redis.MyRedis.SetAccessToken(state, tokenResp.AccessToken, tokenResp.ExpireIn)
	if err != nil {
		return "", err
	}
	return tokenResp.AccessToken, nil
}

// GetUserInformation get user information from UAE Pass
// env environment of UAE Pass
// state to track the user, should be same with before
// token access_token from UAE Pass
func GetUserInformation(env, state, token string) (*GetUserInfoResponse, error) {
	var url string
	if env == "staging" {
		url = config.LocalConfig.Endpoints.Staging.UserInfo
	} else if env == "production" {
		url = config.LocalConfig.Endpoints.Production.UserInfo
	} else {
		return nil, errors.New("the parameter 'env' was invalid")
	}

	//check token availability
	t, err := redis.MyRedis.GetAccessToken(state)
	if err != nil {
		return nil, err
	}
	if token != t {
		return nil, errors.New("the token you passed was invalid")
	}
	resp, err := http_client.Get(url, token, 60)
	if err != nil {
		return nil, err
	}
	//respstr := string(resp)
	//print(respstr)

	var userInfoResp GetUserInfoResponse
	err = json.Unmarshal(resp, &userInfoResp)
	if err != nil {
		return nil, err
	}
	return &userInfoResp, nil
}

// Logout logout UAE Pass
// env environment of UAE Pass
// state to track the user, should be same with before
// token access_token from UAE Pass
// redirect callback from UAE Pass, it's optional
func Logout(env, state, token, redirect string) error {
	var url string
	if env == "staging" {
		url = config.LocalConfig.Endpoints.Staging.Logout
	} else if env == "production" {
		url = config.LocalConfig.Endpoints.Production.Logout
	} else {
		return errors.New("the parameter 'env' was invalid")
	}
	url = fmt.Sprintf("%s?redirect_uri=%s", url, redirect)
	//check token availability
	t, err := redis.MyRedis.GetAccessToken(state)
	if err != nil {
		return err
	}
	if token != t {
		return errors.New("the token you passed was invalid")
	}

	_, err = http_client.Get(url, token, 60)
	if err != nil {
		return err
	}
	//respstr := string(resp)
	//print(respstr)

	return nil
}
