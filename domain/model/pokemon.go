package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Pokemon is the actual representation of the Pokemon objet, contains all  the required details
type Pokemon struct {
	ID   int
	Name string
	Url  string
}

// The string representation of a pokemon
func (p Pokemon) String() string {
	pstr := fmt.Sprintf("%v, %v", p.ID, p.Name)
	return pstr
}

// Creates a pokemon from a string
func ToPokemon(s string) (*Pokemon, error) {
	str := strings.Split(s, ",")
	if len(str) != 3 {
		return &Pokemon{}, errors.New("it's not a pokemon")
	}

	id, err := strconv.ParseInt(str[0], 10, 8)

	if err != nil {
		panic("id is not int")
	}

	return &Pokemon{
		ID:   int(id),
		Name: strings.Trim(str[1], " "),
		Url:  strings.Trim(str[2], " "),
	}, nil
}
