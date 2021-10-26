package main

import (
	"grpc/server/internal/transport/grpc"
	"grpc/server/internal/user"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	userService := user.New()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	go func() {
		grpcClient := grpc.NewGRPC(userService)
		err := grpcClient.Run()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("listening for incoming requests...")

	<-done

	log.Println("finish")
}
