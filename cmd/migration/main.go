package main

import (
	"context"
	"log"

	m "ariga.io/atlas/sql/migrate"

	"go-ent-gqlgen/config"
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/ent/migrate"
	"go-ent-gqlgen/pkg/infrastructure/datastore"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("Failed opening Postgre client: %v", err)
	}

	defer client.Close()

	// createDBFromSchema(client)
	generateMigrationIncrement(client)
}

func createDBFromSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		// migrate.WithGlobalUniqueID(true), // only valid for numeric IDs
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("Failed creating DB schema resouces: %v", err)
	}
}

func generateMigrationIncrement(client *ent.Client) {

	ctx := context.Background()

	// Create or check existence of local migration directory.
	dir, err := m.NewLocalDir("docker/sql/migrations")
	if err != nil {
		log.Fatalf("failed browsing to migration directory: %v", err)
	}

	// Write migration diff.
	err = client.Schema.Diff(ctx, schema.WithDir(dir))
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
