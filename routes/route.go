package routes

import (
	"github.com/avtara/mvc-go/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e:= echo.New()

	e.GET("/client", controller.GetAllClientController)
	e.GET("/client/:id", controller.GetClientByIDController)
	e.POST("/client", controller.PostNewClient)
	e.PUT("/client", controller.PutClient)
	e.DELETE("/client/:id", controller.DeleteClient)

	return e
}