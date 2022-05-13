package registry

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/controller"
	"github.com/GerardoHP/ondemand-go-bootcamp/interactor"
	"github.com/GerardoHP/ondemand-go-bootcamp/presenter"
	"github.com/GerardoHP/ondemand-go-bootcamp/repository"
)

// Retruns a new instance of PokemonController
func (r *registry) NewPokemonController() controller.Pokemon {
	return controller.New(r.NewPokemonInteractor())
}

// Returns a new instance of PokemonInteractor
func (r *registry) NewPokemonInteractor() interactor.Pokemon {
	return interactor.New(r.NewPokemonRepository(), r.NewPokemonPresenter(), nil)
}

// Returns a new instance of PokemonRepository
func (r *registry) NewPokemonRepository() repository.Pokemon {
	return repository.New(r.fileName, nil)
}

// Returns a new instance of PokemonPresenter
func (r *registry) NewPokemonPresenter() presenter.Pokemon {
	return presenter.New()
}
