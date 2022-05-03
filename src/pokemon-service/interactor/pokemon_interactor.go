package interactor

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GerardoHP/ondemand-go-bootcamp/dto"
	"github.com/GerardoHP/ondemand-go-bootcamp/entity"
	"github.com/GerardoHP/ondemand-go-bootcamp/model"
	"github.com/GerardoHP/ondemand-go-bootcamp/presenter"
	"github.com/GerardoHP/ondemand-go-bootcamp/repository"
	"github.com/go-resty/resty/v2"
)

// Implementation of the interface PokemonInteractor
type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

// The interface that holds all the pokemon interactions
type PokemonInteractor interface {
	Get(p []*entity.Pokemon) ([]*entity.Pokemon, error)
	GetPokemon(p string) (*entity.Pokemon, error)
}

// Returns a new PokemonInteractor instance
func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{PokemonRepository: r, PokemonPresenter: p}
}

// Gets all the pokemons found in the repository
func (pk *pokemonInteractor) Get(p []*entity.Pokemon) ([]*entity.Pokemon, error) {
	p, err := pk.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}

	return pk.PokemonPresenter.ResponsePresenter(p), nil
}

// Gets a specific pokemon
func (pk *pokemonInteractor) GetPokemon(p string) (*entity.Pokemon, error) {
	pokemon, err := pk.PokemonRepository.FindByName(p)
	if err != nil {
		return nil, err
	}

	if pokemon == nil {
		newPokemon, _ := getPokemonDetail(p)
		if newPokemon != nil {
			pokemon, _ = pk.PokemonRepository.Add(newPokemon)
		}

	}

	return pokemon, nil
}

// Gets a pokemon detail from a service
func getPokemonDetail(name string) (*entity.Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", strings.ToLower(name))
	client := resty.New()
	resp, err := client.R().EnableTrace().Get(url)
	if err != nil {
		return nil, err
	}

	var pk *dto.Pokemon = &dto.Pokemon{}
	json.Unmarshal(resp.Body(), &pk)
	pk.Url = url
	pk.Image = model.GetImageUrl(pk.ID)

	return &entity.Pokemon{ID: pk.ID, Name: pk.Name, Url: pk.Url, Image: pk.Image}, nil
}
