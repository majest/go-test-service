package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/majest/go-microservice/consul"
	"github.com/majest/go-test-service/handler"
	"github.com/majest/go-test-service/pb"
	"google.golang.org/grpc"
)

var serviceIP, consulIP string
var servicePort, consulPort int

func init() {
	flag.StringVar(&serviceIP, "ip", "127.0.0.1", "Service ip. Should be local ip if run locally")
	flag.IntVar(&servicePort, "port", 9090, "Service port")
	flag.StringVar(&consulIP, "consulip", "", "Consul node ip")
	flag.IntVar(&consulPort, "consulport", 8500, "Consul node port")
	flag.Parse()
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g := grpc.NewServer()
	consul.RegisterService("com.service.string", &consul.Config{
		ServiceIp:   serviceIP,
		ServicePort: servicePort,
		NodeIp:      consulIP,
		NodePort:    consulPort})

	pb.RegisterStringsServer(g, new(handler.StringServer))
	g.Serve(lis)
}
