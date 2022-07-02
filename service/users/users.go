package users

import (
	"context"
	userModel "deall-package/models/users"
	"deall-package/proto"
	types "deall-package/types"
	"fmt"
	"math"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	proto.UnimplementedUsersServiceServer
}

func (s *Server) Create(ctx context.Context, params *proto.UserRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse
	var user types.User
	fmt.Println("params", params)

	// verify empty data
	if params.Email == "" || params.FirstName == "" || params.LastName == "" {
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
	}, 0, 0)

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

func (s *Server) Update(ctx context.Context, params *proto.UserUpdateRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse
	var user types.User
	fmt.Println("params", params)

	// verify empty data
	if params.Id == "" || params.Email == "" || params.FirstName == "" || params.LastName == "" {
		response.Message = "fill required data"
		return &response, nil
	}

	_, err := mail.ParseAddress(params.Email)

	if err != nil {
		response.Message = "invalid email"
		return &response, nil
	}

	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Gender = params.Gender
	user.Email = params.Email
	user.Password = params.Password
	user.Address = params.Address

	_, err = userModel.Update(params.Id, user)

	if err != nil {
		fmt.Println(err)
		response.Message = "Fail"
	} else {
		response.Message = "Success"
	}

	return &response, nil
}

func (s *Server) Delete(ctx context.Context, params *proto.UserDeleteRequest) (*proto.UserResponse, error) {
	var response proto.UserResponse

	_, err := userModel.Delete(params.Id)

	if err != nil {
		fmt.Println(err)
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

func (s *Server) Get(ctx context.Context, params *proto.UserGetRequest) (*proto.UserGetResponse, error) {
	var response proto.UserGetResponse

	paramsBson := bson.M{}
	paramsBson["$or"] = []interface{}{
		bson.M{"email": primitive.Regex{Pattern: params.Search, Options: "i"}},
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.PerPage == 0 {
		params.PerPage = 10
	}

	users := userModel.Find(paramsBson, params.PerPage, (params.PerPage*params.Page)-params.PerPage)
	totalData := userModel.CountDocuments(paramsBson)
	response.TotalPage = int64(math.RoundToEven(float64(totalData) / float64(params.PerPage)))

	for _, val := range users {
		user := proto.UserUpdateRequest{
			Id:        val.ID.String(),
			FirstName: val.FirstName,
			LastName:  val.LastName,
			Gender:    val.Gender,
			Email:     val.Email,
			Address:   val.Address,
		}
		response.Data = append(response.Data, &user)
	}

	response.Message = "Success"
	return &response, nil
}
