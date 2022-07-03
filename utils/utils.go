package utils

import (
	"context"
	"deall-package/types"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Fatal("get metadata error")
	}

	if info.FullMethod != "/proto.UsersService/Login" {
		var claims types.Claims
		token := md["authorization"]

		if len(token) == 0 {
			return "", status.Error(codes.InvalidArgument, "Token not provided")
		}

		tkn, err := jwt.ParseWithClaims(token[0], &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
		})

		if err != nil || !tkn.Valid || claims == (types.Claims{}) {
			return "", status.Error(codes.InvalidArgument, "Token not authorize")
		}

		fmt.Println("claims", claims)
	}

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
