package main

import (
	"fmt"
	"log"
	"net"

	"github.com/majest/go-test-service/consul"
	"github.com/majest/go-test-service/handler"
	"github.com/majest/go-test-service/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g := grpc.NewServer()
	consul.RegisterService("com.service.string")
	pb.RegisterStringsServer(g, new(handler.StringServer))
	g.Serve(lis)
}
