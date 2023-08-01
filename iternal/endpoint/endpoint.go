package endpoint

import (
	"cart_mircoservice/iternal/service"
	"cart_mircoservice/iternal/service/dto"

	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddToCart      endpoint.Endpoint
	DeleteFromCart endpoint.Endpoint
	GetFromCart    endpoint.Endpoint
}

func Make(s service.Service) Endpoints {
	return Endpoints{
		AddToCart:      MakeAddToCart(s),
		DeleteFromCart: MakeDeletefromCart(s),
		GetFromCart:    MakeGetFromCart(s),
	}
}

func MakeAddToCart(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.AddToCartRequest)
		err := s.AddToCart(ctx, req.UserId, req.Item)
		if err != nil {
			return dto.AddToCartResponse{Err: err}, err

		}
		return dto.AddToCartResponse{UserId: req.UserId, Item: req.Item}, nil
	}
}

func MakeDeletefromCart(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteFromCartRequest)
		err := s.DeleteFromCart(ctx, req.UserId, req.ItemId)
		if err != nil {
			return dto.DeleteFromCartResponse{Err: err}, err

		}
		return dto.DeleteFromCartResponse{UserId: req.UserId, ItemId: req.ItemId, Err: err}, nil

	}
}
func MakeGetFromCart(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.GetFromCartRequest)
		cart, err := s.GetFromCart(ctx, req.UserId)
		return dto.GetFromCartResponse{Cart: cart, Err: err}, nil
	}
}
