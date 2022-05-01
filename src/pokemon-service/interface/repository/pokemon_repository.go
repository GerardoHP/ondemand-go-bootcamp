package repository

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
		pk, errPk := toPokemon(scanner.Text())
		if errPk != nil {
			log.Fatal(errPk)
			continue
		}

		p = append(p, pk)
	}

	return p, nil
}

// Creates a pokemon from a string
func toPokemon(s string) (*model.Pokemon, error) {
	str := strings.Split(s, ",")
	if len(str) != 4 {
		return &model.Pokemon{}, errors.New("it's not a pokemon")
	}

	id, err := strconv.ParseInt(str[0], 10, 64)

	if err != nil {
		fmt.Println(s, str[0])
		panic(err)
	}

	return &model.Pokemon{
		ID:    int(id),
		Name:  strings.Trim(str[1], " "),
		Url:   strings.Trim(str[2], " "),
		Image: strings.Trim(str[3], " "),
	}, nil
}
