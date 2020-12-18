package controllers

import (
	"echo-framework/common"
	"echo-framework/helpers"
	"echo-framework/models"
	"net/http"

	"github.com/labstack/echo"
)

//GeneratePassword ...
func GeneratePassword(c echo.Context) (err error) {

	req := new(common.Users)
	if err = c.Bind(req); err != nil {
		return err
	}

	hash, err := helpers.HashPassword(req.Password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) (err error) {

	result, err := models.CheckUser(c)

	return c.JSON(http.StatusOK, result)
}
