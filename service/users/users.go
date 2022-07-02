package users

import (
	"context"
	userModel "deall-package/models/users"
	"deall-package/proto"
	types "deall-package/types"
	"fmt"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
)

type Server struct {
	proto.UnimplementedUsersServiceServer
}

func (s *Server) Create(ctx context.Context, params *proto.UserRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse
	var user types.User
	fmt.Println("params", params)

	// verify empty data
	if params.Email == "" || params.Phone == "" || params.FirstName == "" || params.LastName == "" {
		response.Message = "fill required data"
		return &response, nil
	}

	_, err := mail.ParseAddress(params.Email)

	if err != nil {
		response.Message = "invalid email"
		return &response, nil
	}

	users := userModel.Find(bson.M{
		"email": params.Email,
	})

	if len(users) >= 1 {
		response.Message = "user " + params.Email + " already exist"
		return &response, nil
	}

	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Gender = params.Gender
	user.Email = params.Email
	user.Password = params.Password
	user.Address = params.Address

	_, err = userModel.Insert(user)

	if err != nil {
		response.Message = "Fail"
	} else {
		response.Message = "Success"
	}

	return &response, nil
}

func (s *Server) Login(ctx context.Context, params *proto.UserRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse

	response.Message = "Account not found"

	return &response, nil
}
