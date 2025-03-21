CC = go
INPUT = ./cmd/main.go
OUTPUT = bin/app

build:
	$(CC) build -o $(OUTPUT) $(INPUT)

proto:
	protoc --go_out=. --go-grpc_out=. pkg/proto/nagent.proto 
