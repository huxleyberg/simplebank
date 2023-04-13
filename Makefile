DB_URL=postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgres15 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:15-alpine

create-db:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

drop-db:
	docker exec -it postgres15 dropdb simple_bank

migrate-up:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres create-db drop-db migrate-up migrate-down sqlc