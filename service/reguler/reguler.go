package reguler

import (
	"context"
	regulerModel "deall-package/models/reguler"
	"deall-package/proto"
	"deall-package/types"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Server struct {
	proto.UnimplementedRegulersServiceServer
}

func (s *Server) CreateOrUpdate(ctx context.Context, params *proto.RegulerCreateRequest) (*proto.RegulerResponse, error) {
	var response proto.RegulerResponse
	var regulerData types.Reguler
	customerId, _ := primitive.ObjectIDFromHex(params.CustomerID)
	existingReguler := regulerModel.FindOne(bson.M{
		"customer_id": customerId,
		"tahun":       params.Tahun,
		"bulan":       params.Bulan,
	})

	regulerData.Bulan = params.Bulan
	regulerData.Tahun = params.Tahun
	regulerData.Bina = params.Bina
	regulerData.Infaq = int(params.Infaq)
	regulerData.Zakat = int(params.Zakat)
	regulerData.Shadaqah = int(params.Shadaqah)
	regulerData.Ikhsan = int(params.Ikhsan)
	regulerData.Sabil = int(params.Sabil)
	regulerData.TabunganFitrah = int(params.TabunganFitrah)
	regulerData.TabunganQurban = int(params.TabunganQurban)
	regulerData.CustomerID = customerId

	if existingReguler.ID.IsZero() {
		_, err := regulerModel.Insert(regulerData)

		if err != nil {
			fmt.Println(err)
			response.Message = "Fail"
		} else {
			response.Message = "Success"
		}
	} else {
		result, err := regulerModel.Update(existingReguler.ID, regulerData)
		if err != nil {
			fmt.Println(err)
			response.Message = "Fail"
		} else {
			response.Message = "Success"
		}
		fmt.Println("updateResult", result, regulerData)
	}

	return &response, nil
}

func (s *Server) Update(ctx context.Context, params *proto.RegulerUpdateRequest) (*proto.RegulerResponse, error) {
	var response proto.RegulerResponse

	return &response, nil
}
