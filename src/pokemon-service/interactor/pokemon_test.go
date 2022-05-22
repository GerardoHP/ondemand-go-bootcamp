package interactor

import (
	"errors"
	"net/http"
	"testing"

	"github.com/GerardoHP/ondemand-go-bootcamp/entity"

	"github.com/go-resty/resty/v2"
)

type FakeRepository struct {
	Pokemon         *entity.Pokemon
	AddedPokemon    *entity.Pokemon
	Pokemons        []*entity.Pokemon
	FindByNameError error
	FindAllError    error
	AddError        error
}

func (f FakeRepository) FindAll(p []*entity.Pokemon) ([]*entity.Pokemon, error) {
	return f.Pokemons, f.FindAllError
}

func (f FakeRepository) FindByName(pkName string) (*entity.Pokemon, error) {
	return f.Pokemon, f.FindByNameError
}

func (f FakeRepository) Add(pk *entity.Pokemon) (*entity.Pokemon, error) {
	return f.AddedPokemon, f.AddError
}

func (f FakeRepository) FindAllConcurrent(p []*entity.Pokemon, even bool, n int, n2 int) ([]*entity.Pokemon, error) {
	return nil, nil
}

type FakePresenter struct {
	Pokemons []*entity.Pokemon
	Pokemon  *entity.Pokemon
}

func (f FakePresenter) ResponsePresenter(pk []*entity.Pokemon) []*entity.Pokemon { return f.Pokemons }
func (f FakePresenter) IndividualResponsePresenter(pk *entity.Pokemon) *entity.Pokemon {
	return f.Pokemon
}

type FakeClient struct {
	response      resty.Response
	responseError error
}

func (f FakeClient) Get(url string) (*resty.Response, error) {
	return &f.response, f.responseError
}

// Test to get all of the pokemons
func TestGetPokemons(t *testing.T) {
	repo := FakeRepository{
		Pokemons: []*entity.Pokemon{
			getBasicPokemon(),
		},
	}

	presenter := FakePresenter{
		Pokemons: []*entity.Pokemon{
			getBasicPokemon(),
		},
	}

	client := FakeClient{
		response: resty.Response{},
	}

	interactor := New(repo, presenter, client)
	pks := []*entity.Pokemon{}
	p, err := interactor.Get(pks)
	if err != nil {
		t.Fatal("There shouldn't be an erro")
	}

	if len(p) != 1 {
		t.Fatal("There should only be 1 pokemon")
	}
}

// Test to get all of the pokemons
func TestGetPokemonsErrors(t *testing.T) {
	repo := FakeRepository{
		FindAllError: errors.New("some error"),
	}

	presenter := FakePresenter{}

	client := FakeClient{}

	interactor := New(repo, presenter, client)
	pks := []*entity.Pokemon{}
	_, err := interactor.Get(pks)
	if err == nil {
		t.Fatal("There should be an erro")
	}
}

// Test to get a single pokemon
func TestGetPokemon(t *testing.T) {
	repo := FakeRepository{
		Pokemon: getBasicPokemon(),
	}

	presenter := FakePresenter{
		Pokemon: getBasicPokemon(),
	}

	client := FakeClient{
		response: resty.Response{},
	}

	interactor := New(repo, presenter, client)
	pk := "pokemon"
	p, err := interactor.GetPokemon(pk)
	if err != nil {
		t.Fatal("There shouldn't be an erro")
	}

	if p == nil {
		t.Fatal("There should only be 1 pokemon")
	}
}

// Test to get a single pokemon
func TestGetPokemonDetail(t *testing.T) {
	repo := FakeRepository{
		AddedPokemon: getBasicPokemon(),
	}

	presenter := FakePresenter{
		Pokemon: getBasicPokemon(),
	}

	client := FakeClient{
		response: resty.Response{
			RawResponse: &http.Response{
				StatusCode: http.StatusOK,
			},
		},
	}

	interactor := New(repo, presenter, client)
	pk := "pokemon"
	p, err := interactor.GetPokemon(pk)
	if err != nil {
		t.Fatal("There shouldn't be an erro")
	}

	if p == nil {
		t.Fatal("There should only be 1 pokemon")
	}
}

// Test to get a single pokemon
func TestGetPokemonDetailErrorOnServer(t *testing.T) {
	repo := FakeRepository{
		AddedPokemon: getBasicPokemon(),
	}

	presenter := FakePresenter{
		Pokemon: getBasicPokemon(),
	}

	client := FakeClient{
		responseError: errors.New("some error"),
	}

	interactor := New(repo, presenter, client)
	pk := "pokemon"
	_, err := interactor.GetPokemon(pk)
	if err == nil {
		t.Fatal("There should be an error")
	}

	response := resty.Response{
		RawResponse: &http.Response{
			StatusCode: 404,
		},
	}

	client.responseError = nil
	client.response = response
	interactor = New(repo, presenter, client)
	_, err = interactor.GetPokemon(pk)
	if err == nil {
		t.Fatal("There should be an error")
	}

	client.response.RawResponse.StatusCode = 500
	interactor = New(repo, presenter, client)
	_, err = interactor.GetPokemon(pk)
	if err == nil {
		t.Fatal("There should be an error")
	}

	repo.FindByNameError = errors.New("some error")
	interactor = New(repo, presenter, client)
	_, err = interactor.GetPokemon(pk)
	if err == nil {
		t.Fatal("There should be an error")
	}
}

// Returns a basic default pokemon
func getBasicPokemon() *entity.Pokemon {
	return &entity.Pokemon{ID: 1, Name: "test", Url: "url", Image: "image"}
}
