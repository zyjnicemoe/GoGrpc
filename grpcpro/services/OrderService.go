package services

import (
	"context"
	"log"
)

type OrderService struct {
}

func (this *OrderService) NewOrder(ctx context.Context, request *OrderRequest) (*OrderResponse, error) {
	if err := request.OrderMain.Validate(); err != nil {
		return &OrderResponse{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}
	log.Println(request.OrderMain)
	return &OrderResponse{
		Status:  "OK",
		Message: "success",
	}, nil
}
