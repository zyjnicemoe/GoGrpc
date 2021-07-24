package main

import (
	"google.golang.org/grpc"
	"grpcpro/helper"
	"grpcpro/services"
	"net"
)

func main() {
	//证书"conf/server_no_password.key"
	//creds,err := credentials.NewServerTLSFromFile(crt,key)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	services.RegisterOrderServerServer(rpcServer, new(services.OrderService))
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))
	//tcp
	lis, _ := net.Listen("tcp", ":9029")
	rpcServer.Serve(lis)

	//http
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	//打印使用的协议
	//	log.Println(request)
	//	rpcServer.ServeHTTP(writer, request)
	//})
	//httpServer := http.Server{
	//	Addr:    ":9029",
	//	Handler: mux,
	//}
	//httpServer.ListenAndServeTLS("conf/server/server.pem", "conf/server/server.key")
}
