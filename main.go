package main

import (
	"github.com/majest/go-microservice/server"
	"github.com/majest/go-service-test/pb"
	ss "github.com/majest/go-service-test/server"
	"github.com/majest/go-service-test/service"
)

func main() {

	var svc ss.StringService
	{
		// svc = pureAddService{}
		// svc = loggingMiddleware{svc, logger}
		// svc = instrumentingMiddleware{svc, requestDuration}
	}

	s := server.Init(&server.Config{Name: "Strings Service", Description: "Provides methods to operate on strings"})
	pb.RegisterStringsServer(s.Transport(), service.Binding{svc})
}
