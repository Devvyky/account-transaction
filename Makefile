postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres account_transaction

dropdb:
	docker exec -it postgres dropdb account_transaction

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/account_transaction?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/account_transaction?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server