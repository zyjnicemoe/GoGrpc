package main

import (
	"GoGrpc/grpccli/helper"
	. "GoGrpc/grpccli/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//创建证书
	conn, err := grpc.Dial(":9029", grpc.WithTransportCredentials(helper.GetClientCreds()))
	//grpc.WithInsecure() 不使用证书
	//conn,err := grpc.Dial(":9029",grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	ctx := context.Background()
	//prodClient := NewProdServiceClient(conn)

	//获取单个
	//prodRes, err := prodClient.GetProdStock(ctx,
	//	&ProdRequest{ProdId: 12})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(prodRes.ProdStock)
	//需要加参数GODEBUG=x509ignoreCN=0 go run client.go
	//获取多个
	//response,err:=prodClient.GetProdStocks(ctx,&QuerySize{
	//	Size: 10,
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(response.Prodres[2].ProdStock)
	//枚举测试
	//prodRess, err := prodClient.GetProdStock(ctx,
	//	&ProdRequest{ProdId: 12,ProdArea: ProdAreas_C})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(prodRess.ProdStock)

	//测试商品库存
	//prod,err :=prodClient.GetProdInfo(ctx,&ProdRequest{
	//	ProdId: 12,
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(prod.ProdName)
	//OrderClient := NewOrderServerClient(conn)
	//order,err :=OrderClient.NewOrder(ctx,&OrderRequest{
	//	OrderMain: &OrderMain{
	//	OrderId: 1001,
	//	OrderNo: "2021724",
	//	OrderMoney: 90,
	//	OrderTime: &timestamppb.Timestamp{Seconds: time.Now().Unix()},
	//}})
	//fmt.Println(order.Message)

	//一般模式
	UserClient := NewUserServiceClient(conn)
	var i int32
	req:= UserScoreRequest{}
	req.Users = make([]*UserInfo,0)

	for i = 0; i < 20; i++ {
		req.Users = append(req.Users,&UserInfo{UserId: i})
	}
	res,err := UserClient.GetUserScore(ctx,&req)
	fmt.Println(res.Users)
	//流模式
}
