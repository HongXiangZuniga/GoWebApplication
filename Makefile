install:
	go mod download
build:
	go build -o bin/service cmd/service/service.go
run:
	go run cmd/web/main.go