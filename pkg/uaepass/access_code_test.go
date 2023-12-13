package uaepass

import (
	"UAEPassDemo/pkg/config"
	"github.com/bwmarrin/snowflake"
	"testing"
)

func TestRequestAccessCodeURL(t *testing.T) {
	url := config.LocalConfig.Endpoints.Staging.Authorization
	responseType := "code"
	redirectURL := "http://localhost:8080/receive_code"
	clientID := config.LocalConfig.Endpoints.Staging.ClientID
	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Error(err)
		return
	}
	state := node.Generate().String()
	scope := "urn:uae:digitalid:profile:general"
	acrValues := "urn:safelayer:tws:policies:authentication:level:low"
	language := "en" // en: english, ar: arabic
	accessCodeURL := RequestAccessCodeURL(url, responseType, redirectURL, clientID, state, scope, acrValues, language)
	t.Log("request access code success", accessCodeURL)
}
