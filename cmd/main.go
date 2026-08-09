package main

import (
	"go-explorer/cmd/api"
	"go-explorer/internal/env"
	"log"
)

func main() {
	cfg := api.Config{
		Addr: env.GetEnv("PORT", ":8080"),
	}

	app := &api.Application{
		Config: cfg,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
