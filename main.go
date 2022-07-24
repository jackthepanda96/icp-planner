package main

import (
	"fmt"

	"github.com/jackthepanda96/icp-planner/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	listData []model.User
)

func main() {
	e := echo.New()
	fmt.Println("Hello World")

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.RemoveTrailingSlash())

	// um := model.UserModel{}
	// uc := controller.UserController

	e.Logger.Fatal(e.Start(":80"))
}
