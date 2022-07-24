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

	um := model.UserModel{}
	uc := controller.UserController{Model: um}
	bm := model.BalanceModel{}
	bc := controller.BalanceController{Model: bm}

	e.GET("/users", uc.GetAllUSer())
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	e.PUT("/users/:id", uc.UpdateProfile(), middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if _, err := um.Login(username, password); err != nil {
			return false, nil
		}
		return true, nil
	}))
	e.POST("/income", bc.Income(), middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if _, err := um.Login(username, password); err != nil {
			return false, nil
		}
		return true, nil
	}))
	e.POST("/expense", bc.Expense(), middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if _, err := um.Login(username, password); err != nil {
			return false, nil
		}
		return true, nil
	}))
	e.GET("/balance", bc.Balance(), middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		if _, err := um.Login(username, password); err != nil {
			return false, nil
		}
		return true, nil
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
