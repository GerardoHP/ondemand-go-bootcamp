package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/GerardoHP/ondemand-go-bootcamp/domain/entity"
)

// Pokemon is the actual representation of the Pokemon objet, contains all  the required details
type Pokemon struct {
	entity.Pokemon
}

// Gets the string format of the saved pokemon
func ToString(p entity.Pokemon) string {
	pstr := fmt.Sprintf("%v, %v, %v, %v", p.ID, p.Name, getDetailsUrl(p.ID), GetImageUrl(p.ID))
	return pstr
}

// Creates a pokemon from a string
func ToPokemon(s string) (*entity.Pokemon, error) {
	str := strings.Split(string(s), ",")
	if len(str) != 4 {
		return &entity.Pokemon{}, errors.New("it's not a pokemon")
	}

	id, err := strconv.ParseInt(str[0], 10, 64)

	if err != nil {
		fmt.Println(s, str[0])
		panic(err)
	}

	return &entity.Pokemon{
		ID:    int(id),
		Name:  strings.Trim(str[1], " "),
		Url:   strings.Trim(str[2], " "),
		Image: strings.Trim(str[3], " "),
	}, nil
}

func getDetailsUrl(id int) string {
	return fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", id)
}

func GetImageUrl(id int) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/versions/generation-v/black-white/%v.png", id)
}
