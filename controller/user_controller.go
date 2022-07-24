package controller

import (
	"net/http"

	"github.com/jackthepanda96/icp-planner/model"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.UserModel
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser model.User
		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, "error when parsing data")
		}

		res, err := uc.Model.Insert(newUser)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error from server",
				"status":  false,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success insert user",
			"status":  true,
			"data":    []model.User{res},
		})
	}
}

func (uc *UserController) GetAllUSer() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.Model.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error from server",
				"status":  false,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all user",
			"status":  true,
			"data":    res,
		})
	}
}
