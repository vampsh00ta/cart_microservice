package grpc

import (
	"cart_mircoservice/iternal/service/dto"
	pb "cart_mircoservice/iternal/transport/pb"
	"context"
)

func GRPCEncodeAddToCartResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(dto.AddToCartResponse)
	return &pb.AddToCartResponse{
		UserId: resp.UserId.String(),
		Id:     resp.Item.Id,
		Name:   resp.Item.Name,
		Image:  resp.Item.Image,
		Price:  resp.Item.Price,
	}, nil

}

func GRPCEncodeGetFromCartResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(dto.GetFromCartResponse)
	cart := resp.Cart
	grpcResponse := &pb.GetCartResponse{}
	grpcResponse.Cart = make(map[string]*pb.MapItem)
	for key, value := range cart {
		var mapitem = &pb.MapItem{
			Price: value.Price,
			Image: value.Image,
			Name:  value.Name,
		}
		grpcResponse.Cart[key] = mapitem

	}
	return grpcResponse, nil

}
