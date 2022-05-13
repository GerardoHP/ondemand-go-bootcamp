package dto

type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Url   string `json:"-"`
	Image string `json:"-"`
}
