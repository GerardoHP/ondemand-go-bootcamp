package presenter

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/domain/entity"
)

// Specific  implementation of output for use cases
type PokemonPresenter interface {
	ResponsePresenter(p []*entity.Pokemon) []*entity.Pokemon
	IndividualResponsePresenter(p *entity.Pokemon) *entity.Pokemon
}
