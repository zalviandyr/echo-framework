package routes

import (
	"echo-framework/controllers"

	"github.com/labstack/echo"
)

//CustomersRoute ...
func EmployeesRoute(g *echo.Group) {

	g.GET("/list", controllers.FetchAllEmployees)

	g.POST("/add", controllers.AddEmployees)

	//g.POST("/update", controllers.UpdateCustomer)

	//g.POST("/delete", controllers.DeleteCustomer)

}
