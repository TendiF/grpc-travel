package users

import (
	"context"
	userModel "deall-package/models/users"
	"deall-package/proto"
	types "deall-package/types"
	"deall-package/utils/utils"
	"fmt"
	"log"
	"math"
	"net/mail"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
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
		response.Message = "email user " + params.Email + " already exist"
		return &response, nil
	}

	users = userModel.Find(bson.M{
		"username": params.Username,
	}, 0, 0)

	if len(users) >= 1 {
		response.Message = "username user " + params.Username + " already exist"
		return &response, nil
	}

	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Gender = params.Gender
	user.Email = params.Email
	user.Password = utils.HashAndSalt([]byte(params.Password))
	user.Address = params.Address
	user.Username = params.Username
	user.Role = params.Role.Descriptor().Index()

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

	userFormDB := userModel.FindById(params.Id)

	if userFormDB.ID.Hex() == "000000000000000000000000" {
		response.Message = "user not found"
		return &response, nil
	}

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
	user.Password = utils.HashAndSalt([]byte(params.Password))
	user.Address = params.Address
	user.Username = params.Username

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

	userFormDB := userModel.FindById(params.Id)

	if userFormDB.ID.Hex() == "000000000000000000000000" {
		response.Message = "user not found"
		return &response, nil
	}

	_, err := userModel.Delete(params.Id)

	if err != nil {
		fmt.Println(err)
		response.Message = "Fail"
	} else {
		response.Message = "Success"
	}

	return &response, nil
}

func (s *Server) Login(ctx context.Context, params *proto.UserLoginRequest) (*proto.UserLoginResponse, error) {
	var response proto.UserLoginResponse

	user := userModel.FindOne(bson.M{
		"username": params.Username,
	})

	if utils.ComparePasswords(user.Password, []byte(params.Password)) {
		claims := &types.Claims{
			Uid:  user.ID.Hex(),
			Role: user.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))

		if err != nil {
			response.Message = "fail Signed token"
			return &response, nil
		}
		response.Message = "success"
		response.Token = tokenString
	} else {
		response.Message = "wrong username or password"
	}

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
	response.TotalPage = int64(math.RoundToEven(float64(totalData%params.PerPage)) + 1)

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

func (s *Server) UserProfile(ctx context.Context, params *proto.UserProfileRequest) (*proto.UserProfileResponse, error) {
	var response proto.UserProfileResponse
	var profile proto.UserUpdateRequest
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Fatal("get metadata error")
	}

	if len(md["uid"]) == 0 {
		response.Message = "invalid token uid"
		return &response, nil
	}

	user := userModel.FindById(md["uid"][0])

	profile.FirstName = user.FirstName
	profile.LastName = user.LastName
	profile.Email = user.Email
	profile.Username = user.Username
	profile.Gender = user.Gender
	profile.Role = proto.Role(user.Role)

	response.Profile = &profile

	return &response, nil
}
