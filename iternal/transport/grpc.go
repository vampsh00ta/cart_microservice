package transport

import (
	httpcoding "cart_mircoservice/iternal/coding/grpc"
	"cart_mircoservice/iternal/endpoint"
	pb "cart_mircoservice/iternal/transport/pb"
	"context"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	pb.UnimplementedServiceServer

	addToCart gt.Handler
	getCart   gt.Handler
}

func (g *gRPCServer) AddToCart(ctx context.Context, request *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	_, resp, err := g.addToCart.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.AddToCartResponse), nil
}
func (g *gRPCServer) GetCart(ctx context.Context, request *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	_, resp, err := g.getCart.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetCartResponse), nil
}

func NewGRPCServer(svcEndpoints endpoint.Endpoints, logger log.Logger) pb.ServiceServer {
	return &gRPCServer{
		addToCart: gt.NewServer(
			svcEndpoints.AddToCart,
			httpcoding.GRPCDecodeAddToCartRequest,
			httpcoding.GRPCEncodeAddToCartResponse,
		),
		getCart: gt.NewServer(
			svcEndpoints.GetFromCart,
			httpcoding.GRPCDecodeGetFromCartRequest,
			httpcoding.GRPCEncodeGetFromCartResponse,
		),
	}
}
