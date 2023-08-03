package endpoint

import (
	"cart_mircoservice/iternal/config"
	"cart_mircoservice/iternal/service"
	"cart_mircoservice/iternal/service/dto"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator/v10"
)

type Endpoints struct {
	AddToCart      endpoint.Endpoint
	DeleteFromCart endpoint.Endpoint
	GetFromCart    endpoint.Endpoint
	DeleteCart     endpoint.Endpoint
}
type Middleware func(endpoint.Endpoint) endpoint.Endpoint

//	func JwtMiddleware(logger log.Logger,secretToken string) Middleware {
//		return func(next endpoint.Endpoint) endpoint.Endpoint {
//			return func(ctx context.Context, request interface{}) (interface{}, error) {
//				r := request.(http.Header)
//				authToken := r.Get("Authorization")
//
//				tokenSplited := strings.Split(authToken, " ")
//				if len(tokenSplited) <= 1 {
//					return request,errors.New("CODE_INVALID_AUTH_TOKEN")
//				}
//				rawToken := tokenSplited[1]
//				user :=
//				token, err := jwt.ParseWithClaims(rawToken, user, func(token *jwt.Token) (interface{}, error) {
//
//					return []byte(secretToken), nil
//				})
//				if err != nil {
//					return err
//				}
//				if !token.Valid {
//					return errors.New("CODE_INVALID_AUTH_TOKEN")
//				}
//				return next(ctx, request)
//			}
//		}
//	}
func Make(s service.Service, validate *validator.Validate) Endpoints {
	return Endpoints{
		AddToCart:      MakeAddToCart(s, validate),
		DeleteFromCart: MakeDeletefromCart(s, validate),
		GetFromCart:    MakeGetFromCart(s, validate),
	}
}

func MakeAddToCart(s service.Service, validate *validator.Validate) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.AddToCartRequest)
		fmt.Println(req)
		if err := validate.Struct(&req); err != nil {

			return dto.AddToCartResponse{Err: config.ValidationError}, config.ValidationError
		}
		err := s.AddToCart(ctx, req.UserId, req.Item)
		if err != nil {
			return dto.AddToCartResponse{Err: err}, err

		}
		return dto.AddToCartResponse{UserId: req.UserId, Item: req.Item}, nil
	}
}

func MakeDeletefromCart(s service.Service, validate *validator.Validate) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.DeleteFromCartRequest)
		err := s.DeleteFromCart(ctx, req.UserId, req.ItemId)
		if err != nil {
			return dto.DeleteFromCartResponse{Err: err}, err

		}
		return dto.DeleteFromCartResponse{UserId: req.UserId, ItemId: req.ItemId, Err: err}, nil

	}
}
func MakeGetFromCart(s service.Service, validate *validator.Validate) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.GetFromCartRequest)
		cart, err := s.GetFromCart(ctx, req.UserId)
		return dto.GetFromCartResponse{Cart: cart, Err: err}, nil
	}
}
