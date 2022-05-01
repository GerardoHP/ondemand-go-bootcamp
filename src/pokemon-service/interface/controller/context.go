package controller

// TODO: verify this context
type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
}
