package controller

import (
	"net/http"

	"github.com/jackthepanda96/icp-planner/model"
	"github.com/labstack/echo/v4"
)

type BalanceController struct {
	Model model.BalanceModel
}

func (bc *BalanceController) Income() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.Balance
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "error when parsing data")
		}
		input.Segment = 0
		email, _ := ExtractInfo(c)
		input.Email = email
		res, err := bc.Model.Insert(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error from server",
				"status":  false,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success insert income",
			"status":  true,
			"data":    []model.Balance{res},
		})
	}
}

func (bc *BalanceController) Expense() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.Balance
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "error when parsing data")
		}
		input.Segment = 1
		input.Amount *= -1
		email, _ := ExtractInfo(c)
		input.Email = email
		res, err := bc.Model.Insert(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error from server",
				"status":  false,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success insert expense",
			"status":  true,
			"data":    []model.Balance{res},
		})
	}
}

func (bc *BalanceController) Balance() echo.HandlerFunc {
	return func(c echo.Context) error {
		email, _ := ExtractInfo(c)
		res, total, err := bc.Model.GetBalance(email)
		if total == -1 {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error from server",
				"status":  false,
			})
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error from server",
				"status":  false,
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success insert user",
			"status":  true,
			"data":    res,
			"balance": total,
		})
	}
}
