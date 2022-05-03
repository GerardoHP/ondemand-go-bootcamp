package presenter

import (
	"strings"

	"github.com/GerardoHP/ondemand-go-bootcamp/domain/entity"
)

// The actual implementation of PokemonPresenter interface
type pokemonPresenter struct {
}

// The output of the pokemons, gives the pokemons a special string treatment
type PokemonPresenter interface {
	ResponsePresenter(pk []*entity.Pokemon) []*entity.Pokemon
	IndividualResponsePresenter(pk *entity.Pokemon) *entity.Pokemon
}

// Returns a new instance of PokemonPresenter
func NewPokemonPresenter() PokemonPresenter {
	return &pokemonPresenter{}
}

// Returns a slice pokemons with a special treatment for the pokemon
func (pp *pokemonPresenter) ResponsePresenter(pk []*entity.Pokemon) []*entity.Pokemon {
	for _, p := range pk {
		pokemonTreatment(p)
	}

	return pk
}

// Returs a pokemon with a special treatment
func (pp *pokemonPresenter) IndividualResponsePresenter(pk *entity.Pokemon) *entity.Pokemon {
	pokemonTreatment(pk)
	return pk
}

// Gives a special treatment to the pokemon struct
func pokemonTreatment(p *entity.Pokemon) {
	p.Name = strings.ToUpper(p.Name)
}
