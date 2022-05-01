package presenter

import "github.com/GerardoHP/ondemand-go-bootcamp/domain/model"

// Specific  implementation of output for use cases
type PokemonPresenter interface {
	ResponsePresenter(p []*model.Pokemon) []*model.Pokemon
}
