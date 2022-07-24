package main

import (
	"fmt"
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/pkg/adapter/controller"
	"go-ent-gqlgen/pkg/infrastructure/datastore"
	"go-ent-gqlgen/pkg/infrastructure/graphql"
	"go-ent-gqlgen/pkg/infrastructure/router"
	"go-ent-gqlgen/pkg/registry"
	"log"

	"go-ent-gqlgen/config"
)

const DEFAULT_APP_PORT = "8080"
const DEFAULT_BASE_URL = "http://localhost"

func main() {
	err := config.ReadConfig(config.ReadConfigOption{})

	if err != nil {
		log.Fatalf("Error: Could not load configuration: %v", err)
	}

	dbClient := newDBClient()
	defer dbClient.Close()

	ctrl := newController(dbClient)
	gqlServer := graphql.NewGQLServer(dbClient, ctrl)
	router := router.New(gqlServer)

	fmt.Printf("Server starting on: http://%v", config.C.Server.Address)
	router.Run(fmt.Sprintf("%v", config.C.Server.Address))

}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
