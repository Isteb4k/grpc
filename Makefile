.PHONY: server client

server:
	go run ./server/cmd/server/main.go

client:
	go run ./client/cmd/client/main.go

proto-build:
	protoc --go_out=plugins=grpc:. protos/user/**/*.proto