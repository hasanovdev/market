swag:
	@swag init -g api/main.go -o api/docs

run: swag
	@go run main.go


