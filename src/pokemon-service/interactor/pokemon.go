package interactor

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/GerardoHP/ondemand-go-bootcamp/dto"
	"github.com/GerardoHP/ondemand-go-bootcamp/entity"
	"github.com/GerardoHP/ondemand-go-bootcamp/model"
	"github.com/GerardoHP/ondemand-go-bootcamp/presenter"
	"github.com/GerardoHP/ondemand-go-bootcamp/repository"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	Get(url string) (*resty.Response, error)
}

// Implementation of the interface PokemonInteractor
type pokemonInteractor struct {
	PokemonRepository repository.Pokemon
	PokemonPresenter  presenter.Pokemon
	client            Client
}

// The interface that holds all the pokemon interactions
type Pokemon interface {
	Get(p []*entity.Pokemon) ([]*entity.Pokemon, error)
	GetPokemon(p string) (*entity.Pokemon, error)
	GetEvenOrOdd(p []*entity.Pokemon, even bool, items, items_per_worker int) ([]*entity.Pokemon, error)
}

// Returns a new PokemonInteractor instance
func New(r repository.Pokemon, p presenter.Pokemon, c Client) Pokemon {
	return &pokemonInteractor{PokemonRepository: r, PokemonPresenter: p, client: c}
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
		newPokemon, getPokemonError := getPokemonDetail(p, pk.client)
		if getPokemonError != nil {
			return nil, getPokemonError
		}

		if newPokemon != nil {
			pokemon, _ = pk.PokemonRepository.Add(newPokemon)
		}

	}

	return pokemon, nil
}

// Gets all the pokemons found in the repository that are even or odd
func (pk *pokemonInteractor) GetEvenOrOdd(p []*entity.Pokemon, even bool, items, items_per_worker int) ([]*entity.Pokemon, error) {
	p, err := pk.PokemonRepository.FindAllConcurrent(p, even, items, items_per_worker)
	if err != nil {
		return nil, err
	}

	return pk.PokemonPresenter.ResponsePresenter(p), nil
}

// Gets a pokemon detail from a service
func getPokemonDetail(name string, c Client) (*entity.Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", strings.ToLower(name))
	if c == nil {
		client := resty.New()
		c = client.R().EnableTrace()
	}

	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.RawResponse.StatusCode == http.StatusNotFound {
		return nil, errors.New("Not Found")
	}

	if resp.RawResponse.StatusCode != http.StatusOK {
		return nil, errors.New("Other error")
	}

	var pk *dto.Pokemon = &dto.Pokemon{}
	json.Unmarshal(resp.Body(), &pk)
	pk.Url = url
	pk.Image = model.GetImageUrl(pk.ID)

	return &entity.Pokemon{ID: pk.ID, Name: pk.Name, Url: pk.Url, Image: pk.Image}, nil
}
