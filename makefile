run:
	go run ./cmd/auth/main.go

test:
	go test ./... -v -cover

run-salecar:
	go run ./cmd/salecar/main.go

run-event:
	go run ./cmd/event/main.go

local:
	echo "Starting local environment"
	docker-compose -f docker-compose.local.yml up --build -d

postgres:
	docker run --name salecar_postgres -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:13-alpine

createdb:
	docker exec -it salecar_postgres createdb --username=postgres --owner=postgres salecar

dropdb:
	docker exec -it salecar_postgres dropdb salecar

migrateup:
	migrate -path db/migrations/salecar -database "postgresql://postgres:postgres@localhost:5433/salecar?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations/salecar -database "postgresql://postgres:postgres@localhost:5433/salecar?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown