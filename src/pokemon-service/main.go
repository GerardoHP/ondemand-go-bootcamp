package main

import (
	"fmt"
	"log"

	"github.com/GerardoHP/ondemand-go-bootcamp/config"
	"github.com/GerardoHP/ondemand-go-bootcamp/domain/model"
	"github.com/GerardoHP/ondemand-go-bootcamp/registry"
	"github.com/GerardoHP/ondemand-go-bootcamp/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config := config.GetInstance()
	fn := model.FileName()
	r := registry.NewRegistry(fn)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at port:", config.Port)
	if err := e.Start(":" + config.Port); err != nil {
		log.Fatalln(err)
	}
}
