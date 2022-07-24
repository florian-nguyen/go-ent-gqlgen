package datastore

import (
	"database/sql"
	"fmt"
	"log"

	"go-ent-gqlgen/config"
	"go-ent-gqlgen/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewClient() (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	databaseUrl := GetConnectionString(config.C.Database)
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver)), nil
}

func GetConnectionString(dbConfig config.DatabaseConfig) string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Database)
}
