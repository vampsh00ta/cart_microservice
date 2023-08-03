package main

import (
	pb "cart_mircoservice/iternal/transport/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
)

const (
	jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijg3M2IxZjBmLTZkYzktNDMwZC1iNWRhLTRiMzIwZDFkMTRiZiIsIm5hbWUiOiJvcGl1bSIsImFkbWluIjpmYWxzZSwiZXhwIjoxNjkxMTUyMTMzfQ.T18Nouh6ja6hSfgYLPvce8DBDcaWh7PStO8CT23kT_k"
)

func TestGRPC(t *testing.T) {
	ctx := context.Background()
	//conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//defer conn.Close()
	client := pb.NewServiceClient(conn)
	res1, err := client.AddToCart(ctx,
		&pb.AddToCartRequest{
			Jwt:   jwt,
			Id:    "asdasd",
			Name:  "daun",
			Price: 1234,
			Image: "str",
		})
	res2, err := client.GetCart(ctx,
		&pb.GetCartRequest{
			Jwt: jwt,
		})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", res1)
	log.Printf("Response: %+v", res2)

}
