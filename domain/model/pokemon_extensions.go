package model

import "fmt"

func (p Pokemon) ToString() string {
	pstr := fmt.Sprintf("%v, %v", p.ID, p.Name)
	return pstr
}
