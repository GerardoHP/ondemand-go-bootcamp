package repository

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"github.com/GerardoHP/ondemand-go-bootcamp/entity"
	"github.com/GerardoHP/ondemand-go-bootcamp/model"
)

// The actual implementation of a pokemon repository interface
type pokemonRepository struct {
	pokemonFile string
}

// Repository in charge of all the interactions with the pokemon source file
type PokemonRepository interface {
	FindAll(p []*entity.Pokemon) ([]*entity.Pokemon, error)
	FindByName(pkName string) (*entity.Pokemon, error)
	Add(pk *entity.Pokemon) (*entity.Pokemon, error)
}

// Gets a new instance of pokemon repository
func NewPokemonRepository(fn string) PokemonRepository {
	return &pokemonRepository{pokemonFile: fn}
}

// Gets all the pokemons available
func (repo *pokemonRepository) FindAll(p []*entity.Pokemon) ([]*entity.Pokemon, error) {
	pokemonsMap, err := readAllPokemon(repo.pokemonFile)
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
	pokemonsMap, err := readAllPokemon(repo.pokemonFile)
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
	lines, err := readAllFileLines(repo.pokemonFile)
	if err != nil {
		return nil, err
	}

	fileContent := ""
	pkStr := model.ToString(*pk)
	for _, line := range lines {
		fileContent += line
		fileContent += "\n"
	}

	fileContent += pkStr
	writeErr := ioutil.WriteFile(repo.pokemonFile, []byte(fileContent), 0644)
	if writeErr != nil {
		return nil, writeErr
	}

	return pk, nil
	// file, err := os.Open(repo.pokemonFile)
	// if err != nil {
	// 	log.Fatal("Failed to open", err)
	// 	return nil, err
	// }

	// writer := bufio.NewWriter(file)

	// _, writingError := writer.WriteString(pkStr + "\n")
	// if writingError != nil {
	// 	return nil, writingError
	// }

	// writer.Flush()
	// return pk, nil
}

// Gets all of the pokemon from the file
func readAllPokemon(filename string) (map[string]*model.Pokemon, error) {
	pokemonMap := make(map[string]*model.Pokemon)
	lines, err := readAllFileLines(filename)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		pk, errPk := model.ToPokemon(line)
		if errPk != nil {
			log.Fatal(errPk)
			continue
		}

		pokemonMap[pk.Name] = &model.Pokemon{Pokemon: *pk}
	}

	return pokemonMap, nil
}

func readAllFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open", err)
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
