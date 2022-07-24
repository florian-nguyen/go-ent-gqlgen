package resolver

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/graph/generated"
	"go-ent-gqlgen/pkg/adapter/controller"

	"github.com/99designs/gqlgen/graphql"
)

// Resolver is a context struct
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
		},
	})
}
