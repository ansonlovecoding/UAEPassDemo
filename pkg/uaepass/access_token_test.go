package uaepass

import (
	"UAEPassDemo/pkg/config"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	credential := config.LocalConfig.Endpoints.Staging.Credentials
	token, err := GetAccessToken(
		config.LocalConfig.Endpoints.Staging.Token,
		credential,
		"authorization_code",
		"http://localhost:8080/receive_code",
		"9bf64f8c-3551-362c-9e72-31fdb4f64b3f",
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("token:", token)
}
