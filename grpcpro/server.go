package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcpro/services"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	crt := "conf/server/server.pem"
	key := "conf/server/server.key"
	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}
	certPool := x509.NewCertPool()

	ca, err := ioutil.ReadFile("conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	//证书"conf/server_no_password.key"
	//creds,err := credentials.NewServerTLSFromFile(crt,key)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	//tcp
	//lis, _ := net.Listen("tcp", ":9029")
	//rpcServer.Serve(lis)

	//http
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		rpcServer.ServeHTTP(writer, request)
	})
	httpServer := http.Server{
		Addr:    ":9029",
		Handler: mux,
	}
	httpServer.ListenAndServeTLS(crt, key)
}
