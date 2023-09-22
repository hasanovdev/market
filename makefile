run:
	@go run main.go

swag:
	@swag init -g api/main.go -o api/docs

migrateup:
	@migrate -database postgres://postgres:password@localhost:5432/market\?sslmode\=disable -path migrations up