package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	fileName string = "pokemons.csv"
)

type Pokemon struct {
	ID    int
	Name  string
	Url   string
	Image string
}

func (p Pokemon) String() string {
	pstr := fmt.Sprintf("%v, %v", p.ID, p.Name)
	return pstr
}

func FileName() string { return fileName }

func ToPokemon(s string) (*Pokemon, error) {
	str := strings.Split(s, ",")
	if len(str) != 4 {
		return &Pokemon{}, errors.New("it's not a pokemon")
	}

	id, err := strconv.ParseInt(str[0], 10, 64)

	if err != nil {
		fmt.Println(s, str[0])
		panic(err)
	}

	return &Pokemon{
		ID:    int(id),
		Name:  strings.Trim(str[1], " "),
		Url:   strings.Trim(str[2], " "),
		Image: strings.Trim(str[3], " "),
	}, nil
}
