package controllers

import (
	"echo-framework/models"
	"net/http"

	"github.com/labstack/echo"
)

//FetchAllCustomers ...
func FetchAllCustomers(c echo.Context) (err error) {

	result, err := models.FetchCustomers()

	return c.JSON(http.StatusOK, result)
}

//FetchAllCustomers ...
func StoreCustomer(c echo.Context) (err error) {

	result, err := models.StoreCustomer(c)

	return c.JSON(http.StatusOK, result)
}

//FetchAllCustomers ...
func UpdateCustomer(c echo.Context) (err error) {

	result, err := models.UpdateCustomer(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteCustomer ...
func DeleteCustomer(c echo.Context) (err error) {

	result, err := models.DeleteCustomer(c)

	return c.JSON(http.StatusOK, result)
}
