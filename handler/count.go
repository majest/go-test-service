package handler

import (
	"github.com/majest/go-test-service/pb"
	"golang.org/x/net/context"
	"golang.org/x/net/trace"
)

type Server interface {
	Count(context.Context, *pb.CountRequest) (*pb.CountReply, error)
}

type StringServer struct {
}

func (s *StringServer) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountReply, error) {
	logText(ctx, req.A)
	return &pb.CountReply{V: int64(len(req.A))}, nil
}

func logText(ctx context.Context, value string) {
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("calculating for %s", value)
	}
}
