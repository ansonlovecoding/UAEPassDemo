package uaepass

import (
	"UAEPassDemo/pkg/config"
	"github.com/bwmarrin/snowflake"
	"testing"
)

var accessCode string
var globalState string
var accessToken string

func TestRequestAccessCodeURL(t *testing.T) {
	env := "staging"
	responseType := "code"
	redirectURL := "http://localhost:8080/receive_code"

	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Error(err)
		return
	}
	globalState = node.Generate().String()
	scope := "urn:uae:digitalid:profile:general"
	acrValues := "urn:safelayer:tws:policies:authentication:level:low"
	language := "en" // en: english, ar: arabic
	accessCodeURL, err := RequestAccessCodeURL(env, responseType, redirectURL, globalState, scope, acrValues, language)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("request access code success", accessCodeURL)
}

func TestGetAccessToken(t *testing.T) {
	state := globalState
	grantType := "authorization_code"
	redirect := "http://localhost:8080/receive_code"
	code := accessCode
	token, err := GetAccessToken(
		config.LocalConfig.Env,
		state,
		grantType,
		redirect,
		code,
	)
	if err != nil {
		t.Error(err)
		return
	}
	accessToken = token
	t.Log("token:", token)
}

func TestGetUserInformation(t *testing.T) {
	state := globalState
	token := accessToken
	resp, err := GetUserInformation(
		config.LocalConfig.Env,
		state,
		token)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf(
		"FirstName:%s, LastName:%s, UUID:%s",
		resp.FirstnameEN,
		resp.LastnameEN,
		resp.UUID)
}

func TestLogout(t *testing.T) {
	state := globalState
	token := accessToken
	redirect := "" //optional
	err := Logout(
		config.LocalConfig.Env,
		state,
		token,
		redirect)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Logout UAE Pass success")
}

func TestFlow(t *testing.T) {
	accessCode = "27d66ebe-9193-3993-abb8-175f20b755aa"
	TestGetAccessToken(t)
	TestGetUserInformation(t)
	TestLogout(t)
}
