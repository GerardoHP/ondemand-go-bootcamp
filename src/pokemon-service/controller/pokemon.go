package controller

import (
	"net/http"

	"github.com/GerardoHP/ondemand-go-bootcamp/entity"
	"github.com/GerardoHP/ondemand-go-bootcamp/interactor"
)

// The actual representation of PokemonController interface
type pokemonController struct {
	pokemonInteractor interactor.Pokemon
}

// The controller in charge of getting the pokemosn
type Pokemon interface {
	GetPokemons(c Context) error
	GetPokemon(c Context) error
}

// Returns a new instance a PokemonController
func New(pk interactor.Pokemon) Pokemon {
	return &pokemonController{pokemonInteractor: pk}
}

// Returns all the pokemons from the interactor
func (pc *pokemonController) GetPokemons(c Context) error {
	var p []*entity.Pokemon
	p, err := pc.pokemonInteractor.Get(p)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

// Returns a single pokemon
func (pc *pokemonController) GetPokemon(c Context) error {
	pkName := c.Param("pokemonName")
	p, err := pc.pokemonInteractor.GetPokemon(pkName)
	if err != nil && err.Error() == "Not Found" {
		return c.JSON(http.StatusNotFound, nil)
	}

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
