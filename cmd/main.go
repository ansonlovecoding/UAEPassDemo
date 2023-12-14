package main

import (
	"UAEPassDemo/internal/uaepass"
	"github.com/labstack/echo/v4"
)

func main() {
	c := echo.New()
	c.GET("/receive_code", uaepass.ReceiveCode)
	err := c.Start(":8080")
	if err != nil {
		panic(err)
	}
}
