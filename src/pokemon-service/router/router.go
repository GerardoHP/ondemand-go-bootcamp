package router

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Creates a new router for the application, setting the GET method for the pokemons
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/pokemons", func(context echo.Context) error {
		return c.GetPokemons(context)
	})

	return e
}
