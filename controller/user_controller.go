package controller

import (
	"net/http"
	"strconv"

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

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.User
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "error when parsing data")
		}

		res, err := uc.Model.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "login success",
			"status":  true,
			"data":    res,
		})
	}
}

func (uc *UserController) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.User
		readID := c.Param("id")
		cnv, _ := strconv.Atoi(readID)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "error when parsing data")
		}
		input.ID = cnv

		res, err := uc.Model.Update(input)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update profile",
			"status":  true,
			"data":    res,
		})
	}
}
