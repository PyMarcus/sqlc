postgres:
	docker run --name my_postgres -p 5432:5432 -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it my_postgres createdb --username=myuser --owner=myuser db
dropdb:
	docker exec -it my_postgres dropdb db --username=myuser 
.PHONY: postgres createdb dropdb test run mock

migrateup:
	migrate -path db/migration -database "postgresql://myuser:secret@localhost:5432/db?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://myuser:secret@localhost:5432/db?sslmode=disable" -verbose down
sqlc:
	sqlc generate
start:
	go run main.go 
mock:
	mockgen --destination db/mock/store.go github.com/PyMarcus/go_sqlc/db/sqlc Store

test:
	go test -v -cover ./...
