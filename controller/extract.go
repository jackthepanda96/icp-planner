package controller

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

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
