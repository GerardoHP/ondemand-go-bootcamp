package presenter

import (
	"fmt"

	"github.com/GerardoHP/ondemand-go-bootcamp/domain/model"
)

// Specific  implementation of output for use cases
type PokemonPresenter interface {
	ResponsePresenter(p []*model.Pokemon) []*model.Pokemon
}

// The string representation of a pokemon
func pokemonToString(p model.Pokemon) string {
	pstr := fmt.Sprintf("%v, %v", p.ID, p.Name)
	return pstr
}
