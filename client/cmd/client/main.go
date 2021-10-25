package main

import (
	"context"
	"grpc/client/internal/transport/grpc"
	"grpc/client/internal/user"
	"log"
)

func main() {
	log.Println("start client")

	grpcClient := grpc.NewGRPC()

	userService := user.New(grpcClient)

	ctx := context.Background()

	_, err := userService.CreateUser(ctx, "111", "222")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = userService.GetUser(ctx, "A0000")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = userService.UpdateUser(ctx, "A0000", "111", "222")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("finish")
}
