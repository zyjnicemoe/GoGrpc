package main

import (
	"GoGrpc/grpccli/services"
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair("conf/client/client.pem", "conf/client/client.key")

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
		ServerName:   "zhuyijun",
		RootCAs:      certPool,
	})
	//创建证书
	conn, err := grpc.Dial(":9029", grpc.WithTransportCredentials(creds))
	//grpc.WithInsecure() 不使用证书
	//conn,err := grpc.Dial(":9029",grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId: 12})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(prodRes.ProdStock)
	//需要加参数GODEBUG=x509ignoreCN=0 go run client.go
}
