make schema-up:
	goose -dir internal/db/migration postgres "user=shawntyw password=shawntyw dbname=godb sslmode=disable" up

make schema-down:
	goose -dir internal/db/migration postgres "user=shawntyw password=shawntyw dbname=godb sslmode=disable" down

make query:
	sqlc generate

make run:
	go run main.go