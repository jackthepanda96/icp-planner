package main

import (
	"fmt"

	"github.com/jackthepanda96/icp-planner/controller"
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

	um := model.UserModel{Data: listData}
	uc := controller.UserController{Model: um}

	e.GET("/users", uc.GetAllUSer())
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())

	e.Logger.Fatal(e.Start(":80"))
}
