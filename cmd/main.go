package main

import (
	"UAEPassDemo/internal/uaepass"
	"github.com/labstack/echo/v4"
)

func main() {
	c := echo.New()
	c.GET("/get_access_code_url", uaepass.GetAccessCodeURL)
	c.GET("/receive_code", uaepass.ReceiveCode)
	//c.GET("/get_access_token")
	c.Start(":8080")
}
