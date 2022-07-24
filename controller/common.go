package controller

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/jackthepanda96/icp-planner/model"
	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	ID    int
	Nama  string
	Email string
}

func ParseToResponse(m model.User) UserResponse {
	return UserResponse{
		ID:    m.ID,
		Nama:  m.Nama,
		Email: m.Email,
	}
}

func ParseToResponseArr(m []model.User) []UserResponse {
	var res []UserResponse

	for _, val := range m {
		res = append(res, UserResponse{ID: val.ID, Nama: val.Nama, Email: val.Email})
	}

	return res
}

func ExtractInfo(c echo.Context) (string, error) {
	auth := c.Request().Header.Get("Authorization")
	token := strings.Split(auth, " ")
	parsedCred, err := base64.StdEncoding.DecodeString(token[len(token)-1])
	if err != nil {
		log.Println(err)
		return "", err
	}
	data := strings.Split(string(parsedCred), ":")
	return data[0], nil
}
