createdb:
	docker exec -it postgres createdb --username=root --owner=root db
dropdb:
	docker exec -it postgres dropdb db
entgen:
	go generate ./ent
entdes:
	go run ent describe ./ent/schema
entmig:
	APP_ENV=dev go run ./cmd/migration/main.go
gqlgen:
	gqlgen

.PHONY: createdb dropdb entgen entdes entmig gqlgen