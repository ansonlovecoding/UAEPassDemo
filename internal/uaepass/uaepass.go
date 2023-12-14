package uaepass

import (
	"UAEPassDemo/pkg/redis"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

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
			return c.String(http.StatusOK, fmt.Sprintf("request success, your access code is %s", code))
		}
	} else {
		return c.String(http.StatusBadRequest, "parameters are empty")
	}
}
