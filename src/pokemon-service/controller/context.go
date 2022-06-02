package controller

// Interface to display only required elements of the Echo Context
type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
	QueryParam(name string) string
}
