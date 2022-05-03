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

	e.GET("/pokemon/:pokemonName", func(context echo.Context) error {
		return c.GetPokemon(context)
	})

	//TODO: Add second endpoint to search new pokemons that weren't in the original set,
	// this new endpoint should point to the pokemon api to search for the pokemon and
	// once it has a result add the pokemon to csv file for further consultartionp
	// The API is the following https://pokeapi.co/api/v2/pokemon/{pokemon_name}

	return e
}
