package registry

import (
	"github.com/GerardoHP/ondemand-go-bootcamp/interface/controller"
)

// Implementation of RegistryInterface
type registry struct {
	fileName string
}

// Resolves dependencies using constructor injection
type Registry interface {
	NewAppController() controller.AppController
}

// Returns a new instance of Registry intreface
func NewRegistry(fn string) Registry {
	return &registry{fileName: fn}
}

// Returns a new instance of pokemon controller
func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
