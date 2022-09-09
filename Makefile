install:
	go mod download
build:
	go build -o bin/service cmd/service/main.go
run:
	go run cmd/web/*.go