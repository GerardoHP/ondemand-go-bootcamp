package presenter

import (
	"strings"
	"testing"

	"github.com/GerardoHP/ondemand-go-bootcamp/entity"
)

const (
	pokemonName = "Pikachu"
)

// Tests that the response presenter returns the expected result, a lower case of the pokemon name
func TestResponsePresenter(t *testing.T) {
	pokemons := []*entity.Pokemon{
		getBasicPokemon(),
	}
	presenter := New()
	pokemons = presenter.ResponsePresenter(pokemons)
	if len(pokemons) != 1 {
		t.Fatal("There should be only one pokemon")
	}

	if pokemons[0].Name == strings.ToLower(pokemonName) {
		t.Fatal("The pokemon name should be in lower case")
	}
}

// Tests that the individual response presenter returns the expected result, a lower case of the pokemon name
func TestIndividualPresenter(t *testing.T) {
	pokemon := getBasicPokemon()
	presenter := New()
	pokemon = presenter.IndividualResponsePresenter(pokemon)

	if pokemon.Name == strings.ToLower(pokemonName) {
		t.Fatal("The pokemon name should be in lower case")
	}
}

// Returns a basic structure of pokemon
func getBasicPokemon() *entity.Pokemon {
	return &entity.Pokemon{
		ID:    1,
		Name:  pokemonName,
		Image: "image",
		Url:   "url",
	}
}
