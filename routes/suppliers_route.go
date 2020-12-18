package routes

import (
	"echo-framework/controllers"

	"github.com/labstack/echo"
)

func SuppliersRoute(g *echo.Group) {
	g.GET("/list", controllers.FetchAllSuppliers)

	g.POST("/add", controllers.StoreSupplier)

	g.POST("/update/:supplierID", controllers.UpdateSupplier)

	g.POST("/delete/:supplierID", controllers.DeleteSupplier)
}
