package presenter

import (
	"strings"
	"testing"

	"github.com/GerardoHP/ondemand-go-bootcamp/entity"
)

func TestResponsePresenter(t *testing.T) {
	pokemonName := "Pikachu"
	pokemons := []*entity.Pokemon{
		{ID: 1, Name: pokemonName, Image: "image", Url: "url"},
	}
	presenter := NewPokemonPresenter()
	pokemons = presenter.ResponsePresenter(pokemons)
	if len(pokemons) != 1 {
		t.Fatal("There should be only one pokemon")
	}

	if pokemons[0].Name == strings.ToLower(pokemonName) {
		t.Fatal("The pokemon name should be in lower case")
	}
}

func TestIndividualPresenter(t *testing.T) {
	pokemonName := "Pikachu"
	pokemon := entity.Pokemon{
		ID:    1,
		Name:  pokemonName,
		Image: "image",
		Url:   "url",
	}

	presenter := NewPokemonPresenter()
	pokemon = *presenter.IndividualResponsePresenter(&pokemon)

	if pokemon.Name == strings.ToLower(pokemonName) {
		t.Fatal("The pokemon name should be in lower case")
	}
}
