package registry

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/controller"
	"github.com/GerardoHP/ondemand-go-bootcamp/interactor"

	ip "github.com/GerardoHP/ondemand-go-bootcamp/presenter"
	ir "github.com/GerardoHP/ondemand-go-bootcamp/repository"

	pp "github.com/GerardoHP/ondemand-go-bootcamp/presenter"
	pr "github.com/GerardoHP/ondemand-go-bootcamp/repository"
)

// Retruns a new instance of PokemonController
func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

// Returns a new instance of PokemonInteractor
func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

// Returns a new instance of PokemonRepository
func (r *registry) NewPokemonRepository() pr.PokemonRepository {
	return ir.NewPokemonRepository(r.fileName)
}

// Returns a new instance of PokemonPresenter
func (r *registry) NewPokemonPresenter() pp.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
