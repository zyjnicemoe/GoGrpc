package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpcpro/helper"
	"grpcpro/services"
	"log"
	"net/http"
)

//GODEBUG=x509ignoreCN=0 同样需要加这个参数运行
func main() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}

	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(),
		gwmux,
		"localhost:9029",
		opt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	httpServer := &http.Server{
		Addr:    ":9000",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
