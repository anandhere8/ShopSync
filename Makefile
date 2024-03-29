postgres :
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb :
	docker exec -it postgres12 createdb --username=root --owner=root shopsync

dropdb:
	docker exec -it postgres12 dropdb shopsync

sqlc :
		sqlc generate
test :
		go test -v -cover ./...

migrateup:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shopsync?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shopsync?sslmode=disable" -verbose down

run:
	go run main.go

.PHONY: sqlc test