package users

import (
	"context"

	"deall-package/proto"
)

type Server struct {
	proto.UnimplementedUsersServiceServer
}

func (s *Server) Register(ctx context.Context, params *proto.UserRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse

	// verify empty data
	if params.CountryCode == "" || params.Email == "" || params.Phone == "" || params.FirstName == "" || params.LastName == "" {
		response.Message = "fill required data"
		return &response, nil
	}

	response.Message = "Success"

	return &response, nil
}

func (s *Server) Login(ctx context.Context, params *proto.UserRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse

	response.Message = "Account not found"

	return &response, nil
}
