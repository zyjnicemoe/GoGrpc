package main

import (
	"GoGrpc/grpccli/helper"
	. "GoGrpc/grpccli/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
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

	//一般模式 需要将所有数据传给服务端，等服务端响应
	//UserClient := NewUserServiceClient(conn)
	//var i int32
	//req := UserScoreRequest{}
	//req.Users = make([]*UserInfo, 0)
	//for i = 0; i < 6; i++ {
	//	req.Users = append(req.Users, &UserInfo{UserId: i})
	//}

	//res,err := UserClient.GetUserScore(ctx,&req)
	//fmt.Println(res.Users)
	//流模式 服务端分批返回
	//stream, err := UserClient.GetUserScoreByServerStream(ctx, &req)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for {
	//	res, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	fmt.Println(res.Users)
	//}

	//流模式客户端分批发送
	//userClient := NewUserServiceClient(conn)
	//stream, err := userClient.GetUserScoreByClientStream(ctx)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//var i int32
	//var j int32
	////分三批每次发六条
	//for j = 1; j <= 3; j++ {
	//	req := UserScoreRequest{}
	//	req.Users = make([]*UserInfo, 0)
	//	for i = 1; i < 6; i++ {
	//		//假设是一个耗时任务
	//		req.Users = append(req.Users, &UserInfo{UserId: i*j})
	//	}
	//	if err := stream.Send(&req); err != nil {
	//		log.Fatalln(err)
	//	}
	//}
	//res,_:=stream.CloseAndRecv()
	//fmt.Println(res.Users)

	//双向流
	userClient := NewUserServiceClient(conn)
	stream, err := userClient.GetUserScoreByTWS(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var i int32
	var j int32
	var sum int32 = 0
	//分三批每次发六条
	for j = 1; j <= 3; j++ {
		req := UserScoreRequest{}
		req.Users = make([]*UserInfo, 0)
		for i = 1; i < 6; i++ {
			//假设是一个耗时任务
			sum++
			req.Users = append(req.Users, &UserInfo{UserId: sum})
		}
		if err := stream.Send(&req); err != nil {
			log.Fatalln(err)
		}
		if res, err := stream.Recv(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Println(res.Users)
		}

	}

}
