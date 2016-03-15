package main

import (
	"flag"
	"os"

	"golang.org/x/net/context"

	"github.com/majest/go-microservice/consul"
	"github.com/majest/go-test-service/handler"
	"github.com/majest/go-test-service/pb"
	//"google.golang.org/grpc"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/grpc"
)

var serviceIP, consulIP string
var servicePort, consulPort int

func init() {
	flag.StringVar(&serviceIP, "ip", "127.0.0.1", "Service ip. Should be local ip if run locally")
	flag.IntVar(&servicePort, "port", 9090, "Service port")
	flag.StringVar(&consulIP, "consulip", "192.168.99.101", "Consul node ip")
	flag.IntVar(&consulPort, "consulport", 8500, "Consul node port")
	flag.Parse()
}

func main() {

	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", servicePort))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	logger := log.NewLogfmtLogger(os.Stdout)
	//	g := grpc.NewServer()
	consul.RegisterService("com.service.string", &consul.Config{
		ServiceIp:   serviceIP,
		ServicePort: servicePort,
		NodeIp:      consulIP,
		NodePort:    consulPort})

	//pb.RegisterStringsServer(g, new(handler.StringServer))
	// g.Serve(lis)
	ctx := context.Background()
	s := grpc.NewServer(ctx, makeEndpoint(), DecodeCountRequest, EncodeCountResponse)
	//transportLogger := log.NewContext(logger).With("transport", "gRPC")
	a := &pb.CountRequest{}
	_, r, err := s.ServeGRPC(ctx, a)
	logger.Log("r", r)
	logger.Log("err", err)
}

func makeEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.CountRequest)
		ss := &handler.StringServer{}
		res, err := ss.Count(ctx, &req)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func EncodeCountResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func DecodeCountRequest(ctx context.Context, request interface{}) (interface{}, error) {
	return *request.(*pb.CountRequest), nil
}
