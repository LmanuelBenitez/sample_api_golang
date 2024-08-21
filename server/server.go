package server

import (
	"sample_api/endpoints"

	"github.com/labstack/echo/v4"
)

func InitServer() {
	server := echo.New()
	server.HideBanner = true

	endpoints.SetEndpoints(server)

	server.Logger.Fatal(server.Start(":8080"))
}
