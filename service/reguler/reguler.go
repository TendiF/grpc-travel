package reguler

import (
	"context"
	"deall-package/proto"
)

type Server struct {
	proto.UnimplementedCustomersServiceServer
}

func (s *Server) Create(ctx context.Context, params *proto.RegulerCreateRequest) (*proto.RegulerResponse, error) {
	var response proto.RegulerResponse

	return &response, nil
}

func (s *Server) Update(ctx context.Context, params *proto.RegulerUpdateRequest) (*proto.RegulerResponse, error) {
	var response proto.RegulerResponse

	return &response, nil
}
