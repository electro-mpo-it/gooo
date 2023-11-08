run: docs
	swag init -g ./cmd/app/main.go
	go run ./cmd/app