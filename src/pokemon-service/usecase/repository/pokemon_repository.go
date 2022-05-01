package repository

import "github.com/GerardoHP/ondemand-go-bootcamp/domain/model"

// Repository interface that holds all the operations to interact with the pokemon repository
type PokemonRepositoty interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}
