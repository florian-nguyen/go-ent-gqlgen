package graphql

import (
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/pkg/adapter/controller"
	"go-ent-gqlgen/pkg/adapter/resolver"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
)

// NewServer generates graphql server
func NewGQLServer(client *ent.Client, controller controller.Controller) *handler.Server {
	server := handler.NewDefaultServer(resolver.NewSchema(client, controller))
	server.Use(entgql.Transactioner{TxOpener: client})

	return server
}
