package customers

import (
	"context"
	customerModel "deall-package/models/customer"
	"deall-package/proto"
	"deall-package/types"
	"log"
	"math"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	proto.UnimplementedCustomersServiceServer
}

func (s *Server) Create(ctx context.Context, params *proto.CustomerCreateRequest) (*proto.CustomerResponse, error) {
	var response proto.CustomerResponse
	var customers []interface{}
	var customer types.Customer

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Fatal("get metadata error")
	}
	success := 0
	for _, customerData := range params.Data {

		customer.NIK = customerData.NIK
		customer.Nama = customerData.Nama
		customer.Jk = customerData.Jk
		customer.Status = customerData.Status
		customer.StatusKK = customerData.StatusKK
		customer.NoKK = customerData.NoKK
		customer.KotaKab = customerData.KotaKab
		customer.Kecamatan = customerData.Kecamatan
		customer.DesaKelurahan = customerData.DesaKelurahan
		customer.Kampung = customerData.Kampung
		customer.RT = customerData.RT
		customer.RW = customerData.RW
		customer.Kol = customerData.Kol
		customer.Syahidan = customerData.Syahidan
		customer.TanggalLahir = customerData.TanggalLahir
		customer.PJ = customerData.PJ
		customer.CodeMerchant = md["code_merchant"][0]

		customers = append(customers, customer)
	}

	customerModel.InsertMany(customers)

	response.Message = "success insert : " + strconv.Itoa(success)

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
		var customer proto.Customer
		customer.ID = val.ID.Hex()
		customer.NIK = val.NIK
		customer.Nama = val.Nama
		customer.Jk = val.Jk
		customer.Status = val.Status
		customer.StatusKK = val.StatusKK
		customer.NoKK = val.NoKK
		customer.TanggalLahir = val.TanggalLahir
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
