package repository

import (
	"errors"
	"fmt"
	"testing"

	"github.com/GerardoHP/ondemand-go-bootcamp/entity"
)

type fileUtilsMock struct {
	pokemons         []string
	readAllFiles     error
	appendLineToFile error
}

// Mocks the reading of the files, returning only one file
func (f fileUtilsMock) ReadAllFileLines() ([]string, error) {
	return f.pokemons, f.readAllFiles
}

// Mocks appending a line
func (f fileUtilsMock) AppendLineToFile(line string) error {
	return f.appendLineToFile
}

func (f fileUtilsMock) ReadAllFileConcurrent(even bool, items, items_per_worker int) ([]string, error) {
	return nil, nil
}

// Test that the find all pokemons works as expected
func TestFindAllPokemons(t *testing.T) {
	fakeUtils := fileUtilsMock{
		pokemons: []string{
			"1, pikachu, url, url",
			"some wrong line",
		},
		readAllFiles:     nil,
		appendLineToFile: nil,
	}

	repo := New("filelocation", fakeUtils)

	var pk []*entity.Pokemon
	pk, err := repo.FindAll(pk)

	if err != nil {
		t.Fatal("There shouldn't be any error")
	}

	if len(pk) != 1 {
		t.Fatal("There should be only one pokemon")
	}
}

// Test mocking an errror during the read should return an error
func TestReadingErrors(t *testing.T) {
	fakeUtils := fileUtilsMock{
		pokemons:         []string{},
		readAllFiles:     errors.New("Random error"),
		appendLineToFile: errors.New("Random error"),
	}
	repo := New("filelocation", fakeUtils)

	var pk []*entity.Pokemon
	_, err := repo.FindAll(pk)
	if err == nil {
		t.Fatal("There should be an error finding all pokemons")
	}

	_, err = repo.FindByName("any")
	if err == nil {
		t.Fatal("There should be an error finding a pokemon")
	}

	p := entity.Pokemon{
		ID:    0,
		Name:  "",
		Url:   "",
		Image: "",
	}

	_, err = repo.Add(&p)
	if err == nil {
		t.Fatal("There should be an error adding a pokemon on reading pokemons")
	}

	fakeUtils.readAllFiles = nil
	repo = New("filelocation", fakeUtils)
	_, err = repo.Add(&p)
	if err == nil {
		t.Fatal("There should be an error adding a pokemon on appending a file ")
	}
}

func TestFindByName(t *testing.T) {
	fakeUtils := fileUtilsMock{
		pokemons:         []string{"1, pikachu, url, url"},
		readAllFiles:     nil,
		appendLineToFile: nil,
	}
	repo := New("filelocation", fakeUtils)

	pk, _ := repo.FindByName("pikachu")
	if pk == nil {
		t.Fatal("There should be a pikachu pokemon")
	}

	pk, _ = repo.FindByName("raichu")
	fmt.Println(pk)
	if pk != nil {
		t.Fatal("There shouldn't be any raichu pokemon")
	}
}

func TestAddingPokemon(t *testing.T) {
	fakeUtils := fileUtilsMock{
		pokemons:         []string{"1, pikachu, url, url"},
		readAllFiles:     nil,
		appendLineToFile: nil,
	}
	repo := New("location", fakeUtils)

	p := &entity.Pokemon{
		ID:    1,
		Name:  "name",
		Url:   "url",
		Image: "image",
	}
	p, err := repo.Add(p)
	if err != nil {
		t.Fatal("There shouldn't be any error on adding a new pokemon")
	}

	p.Name = "pikachu"
	_, err = repo.Add(p)
	if err == nil {
		t.Fatal("There should't be able to add again pikachu pokemon")
	}
}
