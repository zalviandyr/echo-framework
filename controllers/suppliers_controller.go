package controllers

import (
	"echo-framework/models"
	"net/http"

	"github.com/labstack/echo"
)

func FetchAllSuppliers(c echo.Context) (err error) {
	result, err := models.FetchSuppliers()

	return c.JSON(http.StatusOK, result)
}

func StoreSupplier(c echo.Context) (err error) {
	result, err := models.StoreSupplier(c)

	return c.JSON(result.Status, result)
}

func UpdateSupplier(c echo.Context) (err error) {
	result, err := models.UpdateSupplier(c)

	return c.JSON(result.Status, result)
}

func DeleteSupplier(c echo.Context) (err error) {
	result, err := models.DeleteSupplier(c)

	return c.JSON(result.Status, result)
}
