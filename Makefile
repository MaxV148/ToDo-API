postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root check_todo
dropdb:
	docker exec -it postgres16 dropdb check_todo

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/check_todo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/check_todo?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover sqlc

server:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown sqlc test server