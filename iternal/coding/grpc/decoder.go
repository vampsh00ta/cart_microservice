package grpc

import (
	"cart_mircoservice/iternal/coding"
	"cart_mircoservice/iternal/config"
	"cart_mircoservice/iternal/db/redis"
	"cart_mircoservice/iternal/service/dto"
	pb "cart_mircoservice/iternal/transport/pb"
	"context"
)

func GRPCDecodeAddToCartRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddToCartRequest)

	var user config.JwtCustomClaim
	if err := coding.DecodeJwt("secret", &user, req.Jwt); err != nil {
		return req, err
	}
	item := redis.Item{
		req.Id,
		redis.MapItem{
			Image: req.Image,
			Price: req.Price,
			Name:  req.Name,
		},
	}
	return dto.AddToCartRequest{user.Id, item}, nil
	//return dto.GRPCAddToCartRequest{JWT: req.Jwt,Item:{}}, nil
}

func GRPCDecodeGetFromCartRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetCartRequest)

	var user config.JwtCustomClaim
	if err := coding.DecodeJwt("secret", &user, req.Jwt); err != nil {
		return req, err
	}

	return dto.GetFromCartRequest{UserId: user.Id}, nil
	//return dto.GRPCAddToCartRequest{JWT: req.Jwt,Item:{}}, nil
}
