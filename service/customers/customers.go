package customers

import (
	"context"
	customerModel "deall-package/models/customer"
	userModel "deall-package/models/users"
	"deall-package/proto"
	"deall-package/types"
	"fmt"
	"log"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	proto.UnimplementedCustomersServiceServer
}

func (s *Server) Create(ctx context.Context, params *proto.CustomerCreateRequest) (*proto.CustomerResponse, error) {
	var response proto.CustomerResponse
	var customer types.Customers

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal("get metadata error")
	}
	user := userModel.FindById(md["uid"][0])

	customer.No = params.No
	customer.NIK = params.NIK
	customer.Nama = params.Nama
	customer.StatusKK = params.StatusKK
	customer.NoKK = params.NoKK
	customer.KotaKab = params.KotaKab
	customer.Kecamatan = params.Kecamatan
	customer.DesaKelurahan = params.DesaKelurahan
	customer.Kampung = params.Kampung
	customer.RT = params.RT
	customer.RW = params.RW
	customer.Kol = params.Kol
	customer.Syahidan = params.Syahidan
	customer.PJ = params.PJ
	customer.CodeMerchant = user.CodeMerchant

	createdCustomer, err := customerModel.Insert(customer)

	if err != nil {
		response.Message = "fail"
	} else {
		response.Message = "success"
	}

	fmt.Println(createdCustomer)

	return &response, nil
}

func (s *Server) Get(ctx context.Context, params *proto.CustomerGetRequest) (*proto.CustomerGetResponse, error) {
	var response proto.CustomerGetResponse

	paramsBson := bson.M{}
	paramsBson["$or"] = []interface{}{
		bson.M{"nama": primitive.Regex{Pattern: params.Search, Options: "i"}},
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.PerPage == 0 {
		params.PerPage = 10
	}

	customers := customerModel.Find(paramsBson, params.PerPage, (params.PerPage*params.Page)-params.PerPage)
	totalData := customerModel.CountDocuments(paramsBson)
	response.TotalPage = int64(math.RoundToEven(float64(totalData%params.PerPage)) + 1)

	for _, val := range customers {
		var customer proto.CustomerCreateRequest
		customer.No = val.No
		customer.NIK = val.NIK
		customer.Nama = val.Nama
		customer.StatusKK = val.StatusKK
		customer.NoKK = val.NoKK
		customer.TanggalLahir = val.TanggalLahir
		customer.Usia = val.Usia
		customer.KotaKab = val.KotaKab
		customer.Kecamatan = val.Kecamatan
		customer.DesaKelurahan = val.DesaKelurahan
		customer.Kampung = val.Kampung
		customer.RT = val.RT
		customer.RW = val.RW
		customer.Kol = val.Kol
		customer.Syahidan = val.Syahidan
		customer.PJ = val.PJ
		response.Data = append(response.Data, &customer)
	}

	response.Message = "success"

	return &response, nil
}
