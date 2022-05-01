package repository

import (
	"bufio"
	"log"
	"os"

	"github.com/GerardoHP/ondemand-go-bootcamp/domain/model"
)

// The actual implementation of a pokemon repository interface
type pokemonRepository struct {
	pokemonFile string
}

// Repository in charge of all the interactions with the pokemon source file
type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}

// Gets a new instance of pokemon repository
func NewPokemonRepository(fn string) PokemonRepository {
	return &pokemonRepository{pokemonFile: fn}
}

// Gets all the pokemons available
func (pk *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	file, err := os.Open(pk.pokemonFile)
	if err != nil {
		log.Fatal("Failed to open", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		pk, errPk := model.ToPokemon(scanner.Text())
		if errPk != nil {
			log.Fatal(err)
			continue
		}

		p = append(p, pk)
	}

	return p, nil
}
