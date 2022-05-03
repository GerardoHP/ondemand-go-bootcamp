package repository

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/domain/entity"
)

// Repository interface that holds all the operations to interact with the pokemon repository
type PokemonRepositoty interface {
	FindAll(p []*entity.Pokemon) ([]*entity.Pokemon, error)
	FindByName(p string) (*entity.Pokemon, error)
	Add(pk *entity.Pokemon) (*entity.Pokemon, error)
}
