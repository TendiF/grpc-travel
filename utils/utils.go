package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

func UnaryInterceptorHandler(ctx context.Context,
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

	CreateConnection(3 * time.Second)

	// Calls the handler
	h, err := handler(ctx, req)

	return h, err
}

func WithServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(UnaryInterceptorHandler)
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
