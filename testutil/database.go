package testutil

import (
	"context"
	"database/sql"
	"go-ent-gqlgen/config"
	"go-ent-gqlgen/ent"
	"go-ent-gqlgen/pkg/infrastructure/datastore"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4"
)

// NewDBClient loads database for test.
func NewDBClient(t *testing.T) *ent.Client {
	connString := datastore.GetConnectionString(config.C.Database)
	db, err := sql.Open("pgx", connString)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}

// DropAll drops all data from database.
func DropAll(t *testing.T, client *ent.Client) {
	t.Log("drop data from database")
	DropUser(t, client)
	DropTodo(t, client)
}

// DropUser drops data from users.
func DropUser(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.User.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// DropTodo drops data from todos.
func DropTodo(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	_, err := client.Todo.Delete().Exec(ctx)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
