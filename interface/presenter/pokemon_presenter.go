package presenter

import (
	"strings"

	"github.com/GerardoHP/ondemand-go-bootcamp/domain/model"
)

// The actual implementation of PokemonPresenter interface
type pokemonPresenter struct {
}

// The output of the pokemons, gives the pokemons a special string treatment
type PokemonPresenter interface {
	ResponsePresenter(pk []*model.Pokemon) []*model.Pokemon
}

// Returns a new instance of PokemonPresenter
func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

// Returns a slice pokemons,
func (pp *pokemonPresenter) ResponsePresenter(pk []*model.Pokemon) []*model.Pokemon {
	for _, p := range pk {
		p.Name = strings.ToUpper(p.Name)
	}

	return pk
}
