postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine
createdb:
	docker exec -it postgres16 createdb --username=root --owner=root check_todo

createdb-prod:
	docker exec -it postgres16 createdb --username=root --owner=root check_todo_prod

dropdb:
	docker exec -it postgres16 dropdb check_todo

dropdb-prod:
	docker exec -it postgres16 dropdb check_todo_prod

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/check_todo?sslmode=disable" -verbose up

migrateup-prod:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/check_todo_prod?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/check_todo?sslmode=disable" -verbose down

migratedown-prod:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/check_todo_prod?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover sqlc

server:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown sqlc test server createdb-prod dropdb-prod migratedown-prod migrateup-prod