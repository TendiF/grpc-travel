package utils

import (
	"context"
	userModel "deall-package/models/users"
	"deall-package/types"
	"deall-package/utils/database"
	"strconv"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	database.CreateConnection(3 * time.Second)

	if info.FullMethod != "/proto.UsersService/Login" {
		var claims types.Claims
		token := md["authorization"]

		if len(token) == 0 {
			return "", status.Error(codes.Unauthenticated, "Token not provided")
		}

		tkn, err := jwt.ParseWithClaims(token[0], &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SIGNATURE_KEY")), nil
		})

		if err != nil || !tkn.Valid || claims == (types.Claims{}) {
			return "", status.Error(codes.Unauthenticated, "Token not authorize")
		}

		endpoint := map[string]string{
			"/proto.UsersService/Create": "0",
			"/proto.UsersService/Get":    "0",
			"/proto.UsersService/Update": "0",
			"/proto.UsersService/Delete": "0",
		}

		if endpoint[info.FullMethod] != "" && endpoint[info.FullMethod] != strconv.Itoa(claims.Role) {
			return nil, status.Error(codes.PermissionDenied, "Don't have access")
		}

		md.Append("uid", claims.Uid)
		user := userModel.FindById(claims.Uid)

		if user == (types.User{}) {
			return nil, status.Error(codes.PermissionDenied, "User not found "+claims.Uid)
		}

		md.Append("code_merchant", user.CodeMerchant)

		if user.CodeMerchant == "" {
			return nil, status.Error(codes.PermissionDenied, "Code merchant not set")
		}

		ctx = metadata.NewIncomingContext(ctx, md)
	}

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

func InitAdmin() {
	var user types.User

	user = userModel.FindByUsername("admin")

	if user.ID == (primitive.ObjectID{}) {
		user.Username = "admin"
		user.Email = "admin@admin.com"
		user.Role = 0
		user.Password = HashAndSalt([]byte("admin"))
		user.CodeMerchant = "35"

		userModel.Insert(user)
	}

	user = userModel.FindByUsername("user")

	if user.ID == (primitive.ObjectID{}) {
		user.Username = "user"
		user.Email = "user@user.com"
		user.Role = 1
		user.Password = HashAndSalt([]byte("user"))
		user.CodeMerchant = "35"

		userModel.Insert(user)
	}
}
