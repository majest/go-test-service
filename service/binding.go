package service

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/majest/go-service-test/pb"
	"github.com/majest/go-service-test/server"
	"golang.org/x/net/context"
)

type Binding struct {
	server.StringService
}

func (b Binding) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountReply, error) {
	return &pb.CountReply{V: int64(b.StringService.Count(req.A))}, nil
}

func CountEndpoint(svc server.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.CountRequest)
		return &pb.CountReply{V: int64(svc.Count(req.A))}, nil
	}
}
