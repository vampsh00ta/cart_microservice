package transport

import (
	httpcoding "cart_mircoservice/iternal/coding/http"
	"cart_mircoservice/iternal/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewService(
	svcEndpoints endpoint.Endpoints, logger log.Logger,
) http.Handler {
	// set-up router and initialize http endpoints
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(httpcoding.EncodeErrorResponse),
	}
	// HTTP Post - /orders
	r.Methods("GET").Path("/items/").Handler(kithttp.NewServer(
		svcEndpoints.GetFromCart,
		httpcoding.DecodeGetFromCart,
		httpcoding.EncodeResponse,
		options...,
	))

	// HTTP Post - /orders/{id}
	r.Methods("POST").Path("/items/").Handler(kithttp.NewServer(
		svcEndpoints.AddToCart,
		httpcoding.DecodeAddToCartRequest,
		httpcoding.EncodeResponse,
		options...,
	))

	// HTTP Post - /orders/status
	r.Methods("DELETE").Path("/items/").Handler(kithttp.NewServer(
		svcEndpoints.DeleteFromCart,
		httpcoding.DecodeDeleteFromCartRequest,
		httpcoding.EncodeResponse,
		options...,
	))

	return r
}
