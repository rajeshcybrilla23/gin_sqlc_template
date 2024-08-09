DB_URL=postgres://rajeshmanjunath:password@localhost:5432/gin_sqlc_template?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockDB -destination db/mock/store.go  github.com/raj23manj/techshool-go/db/sqlc Store

.PHONY: migrateup migratedown migrateup1 migratedown1 new_migration sqlc test server mock
