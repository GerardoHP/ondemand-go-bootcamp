package repository

import (
	"errors"
	"log"

	"github.com/GerardoHP/ondemand-go-bootcamp/entity"
	"github.com/GerardoHP/ondemand-go-bootcamp/model"
	"github.com/GerardoHP/ondemand-go-bootcamp/utils"
)

// The actual implementation of a pokemon repository interface
type pokemonRepository struct {
	pokemonFile string
	fileUtils   utils.File
}

// Repository in charge of all the interactions with the pokemon source file
type Pokemon interface {
	FindAll(p []*entity.Pokemon) ([]*entity.Pokemon, error)
	FindByName(pkName string) (*entity.Pokemon, error)
	Add(pk *entity.Pokemon) (*entity.Pokemon, error)
	FindAllConcurrent(p []*entity.Pokemon, n int, even bool) ([]*entity.Pokemon, error)
}

// Gets a new instance of pokemon repository
func New(fn string, f utils.File) Pokemon {
	if f == nil {
		f = utils.New(fn)
	}

	return &pokemonRepository{pokemonFile: fn, fileUtils: f}
}

// Gets all the pokemons available
func (repo *pokemonRepository) FindAll(p []*entity.Pokemon) ([]*entity.Pokemon, error) {
	pokemonsMap, err := readAllPokemon(repo.fileUtils)
	if err != nil {
		return nil, err
	}

	for _, pokemon := range pokemonsMap {
		p = append(p, &pokemon.Pokemon)
	}

	return p, nil
}

// Gets a pokemon from the repository, returns nil in case it's not found
func (repo *pokemonRepository) FindByName(pkName string) (*entity.Pokemon, error) {
	pokemonsMap, err := readAllPokemon(repo.fileUtils)
	if err != nil {
		return nil, err
	}

	pokemonModel := pokemonsMap[pkName]
	if pokemonModel == nil {
		return nil, nil
	}

	return &pokemonModel.Pokemon, nil
}

// Adds a pokemon to the file
func (repo *pokemonRepository) Add(pk *entity.Pokemon) (*entity.Pokemon, error) {
	pkD, err := repo.FindByName(pk.Name)
	if pkD != nil {
		return nil, errors.New("duplicated pokemon")
	}

	if err != nil {
		return nil, err
	}

	err = repo.fileUtils.AppendLineToFile(model.ToString(*pk))
	if err != nil {
		return nil, err
	}

	return pk, nil
}

func (repo *pokemonRepository) FindAllConcurrent(p []*entity.Pokemon, n int, even bool) ([]*entity.Pokemon, error) {
	pokemonsMap, err := readAllPokemonConcurrent(n, even, repo.fileUtils)
	if err != nil {
		return nil, err
	}

	for _, pokemon := range pokemonsMap {
		p = append(p, &pokemon.Pokemon)
	}

	return p, nil
}

// Gets all of the pokemon from the file
func linesToPokemons(lines []string) (map[string]*model.Pokemon, error) {
	pokemonMap := make(map[string]*model.Pokemon)
	for _, line := range lines {
		pk, errPk := model.ToPokemon(line)
		if errPk != nil {
			log.Println(errPk)
			continue
		}

		pokemonMap[pk.Name] = &model.Pokemon{Pokemon: *pk}
	}

	return pokemonMap, nil
}

func readAllPokemonConcurrent(n int, even bool, fUtils utils.File) (map[string]*model.Pokemon, error) {
	lines, err := fUtils.ReadAllFileConcurrent(n, even)
	if err != nil {
		return nil, err
	}

	return linesToPokemons(lines)
}

func readAllPokemon(fUtils utils.File) (map[string]*model.Pokemon, error) {
	lines, err := fUtils.ReadAllFileLines()
	if err != nil {
		return nil, err
	}

	return linesToPokemons(lines)
}
