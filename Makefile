postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres account_transaction

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/account_transaction?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/account_transaction?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/devvyky/account-transaction/db/sqlc Store

.PHONY: postgres createdb migrateup migratedown sqlc test server mock