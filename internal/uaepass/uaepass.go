package uaepass

import (
	"UAEPassDemo/pkg/config"
	"UAEPassDemo/pkg/redis"
	"UAEPassDemo/pkg/uaepass"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetAccessCodeURL get the url to login UAE Pass
func GetAccessCodeURL(c echo.Context) error {
	var url string
	var clientID string

	redirect := c.QueryParam("redirect")
	env := c.QueryParam("env")
	if env == "staging" {
		url = config.LocalConfig.Endpoints.Staging.Authorization
		clientID = config.LocalConfig.Endpoints.Staging.ClientID
	} else {
		url = config.LocalConfig.Endpoints.Production.Authorization
		clientID = config.LocalConfig.Endpoints.Production.ClientID
	}

	responseType := "code"
	node, err := snowflake.NewNode(1)
	if err != nil {
		c.String(http.StatusBadRequest, "failed in generating state")
		return err
	}
	state := node.Generate().String()
	scope := "urn:uae:digitalid:profile:general"
	acrValues := "urn:safelayer:tws:policies:authentication:level:low"
	language := "en" // en: english, ar: arabic
	accessCodeURL := uaepass.RequestAccessCodeURL(url, responseType, redirect, clientID, state, scope, acrValues, language)
	c.String(http.StatusOK, accessCodeURL)
	return nil
}

// ReceiveCode Receive parameters "state" and "code" from the call back
func ReceiveCode(c echo.Context) error {
	state := c.QueryParam("state")
	code := c.QueryParam("code")

	// store the access code to redis with expiration time
	if len(state) > 0 && len(code) > 0 {
		//expiration time of access code is 10 minutes
		err := redis.MyRedis.SetAccessCode(state, code, 10*60)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		} else {
			return c.String(http.StatusOK, "request success")
		}
	} else {
		return c.String(http.StatusBadRequest, "parameters are empty")
	}
}

func GetAccessToken(c echo.Context) error {
	state := c.QueryParam("state")
	code := c.QueryParam("code")

	//compare with the redis
	if len(state) > 0 && len(code) > 0 {
		code2, err := redis.MyRedis.GetAccessCode(state)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return err
		}
		if code != code2 {
			c.String(http.StatusBadRequest, "code is not for this state")
			return errors.New("code is not for this state")
		}

		return nil
	} else {
		return c.String(http.StatusBadRequest, "parameters are empty")
	}

}
