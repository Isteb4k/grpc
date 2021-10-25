package main

import (
	"grpc/server/internal/transport/grpc"
	"grpc/server/internal/user"
	"log"
)

func main() {
	userService := user.New()

	grpcClient := grpc.NewGRPC(userService)
	err := grpcClient.Run()
	if err != nil {
		log.Fatal(err)
	}
}
