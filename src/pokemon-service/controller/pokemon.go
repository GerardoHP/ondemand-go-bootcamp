package controller

import (
	"errors"
	"net/http"
	"strconv"

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
	GetPokemonsEvenOrOdd(c Context) error
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

// Returns all the pokemons from the interactor
func (pc *pokemonController) GetPokemonsEvenOrOdd(c Context) error {
	var p []*entity.Pokemon

	even, err := getIsEvenOrOdd(c.QueryParam("type"))
	if err != nil {
		return err
	}

	items, err := strconv.Atoi(c.QueryParam("items"))
	if err != nil {
		return err
	}

	items_per_worker, err := strconv.Atoi(c.QueryParam("items_per_worker"))
	if err != nil {
		return err
	}

	p, err = pc.pokemonInteractor.GetEvenOrOdd(p, even, items, items_per_worker)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

// Gets if a string is even or odd
func getIsEvenOrOdd(str string) (bool, error) {
	switch str {
	case "even":
		return true, nil
	case "odd":
		return false, nil
	default:
		return false, errors.New("string not supported")
	}
}
