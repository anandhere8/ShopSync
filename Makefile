sqlc :
		sqlc generate

test :
		go test -v -cover ./...

migrateup:
		migrate -path db/schema/ShopSync.sql -database "postgresql://root:secret@localhost:5432/shopsync?sslmode=disable" -verbose up

.PHONY: sqlc test