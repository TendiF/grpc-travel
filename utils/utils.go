package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
)

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	s := fmt.Sprintf("%s|%s", info.FullMethod, req)

	file, err := os.OpenFile("./tmp/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("ACCESS|" + s)

	// Calls the handler
	h, err := handler(ctx, req)

	return h, err
}

func WithServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}
