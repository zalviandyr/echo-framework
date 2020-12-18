package routes

import (
	"echo-framework/controllers"

	"github.com/labstack/echo"
)

//UserRoute is ...
func UserRoute(g *echo.Group) {

	g.GET("/list", controllers.FetchAllUsers)

	g.POST("/add", controllers.StoreUser)

	g.POST("/update", controllers.UpdateUser)

	g.POST("/delete", controllers.DeleteUser)

	g.POST("/login", controllers.CheckLogin)

}
