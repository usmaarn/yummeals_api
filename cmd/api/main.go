package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/usmaarn/yummeals_api/cmd/application"
)

func main() {
	r := chi.NewRouter()

	app := application.New(r)

	app.RegisterMiddleware()
	app.RegisterRoutes()

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
