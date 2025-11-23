package main

import (
	"go-circleci/api"
	"go-circleci/logger"
	"go-circleci/services"
	"log"
)

func main() {
	service := services.NewCatFactService("https://catfact.ninja/fact")

	service = logger.NewLoggingService(service)

	// fact, err := service.GetCatFact(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v\n", fact)

	apiServer := api.NewApiServer(service)
	log.Fatal(apiServer.Start(":3000"))
}
