package services

import (
	"context"
)

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch request.ProdArea {
	case ProdAreas_A:
		stock = 30
	case ProdAreas_B:
		stock = 31
	case ProdAreas_C:
		stock = 50
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, request *QuerySize) (*ProdResponseList, error) {
	Prodres := []*ProdResponse{
		{ProdStock: 20},
		{ProdStock: 30},
		{ProdStock: 50},
		{ProdStock: 90},
	}
	return &ProdResponseList{
		Prodres: Prodres,
	}, nil
}
func (this *ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
	return &ProdModel{ProdId: 101,
		ProdName:  "测试商品",
		ProdPrice: 20.5,
	}, nil
}
