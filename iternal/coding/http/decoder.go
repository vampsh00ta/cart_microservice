package http

import (
	"cart_mircoservice/iternal/config"
	"cart_mircoservice/iternal/service/dto"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func DecodeGetFromCart(_ context.Context, r *http.Request) (request interface{}, err error) {
	var user config.JwtCustomClaim

	if err := DecodeJwtHttp(r, "secret", &user); err != nil {
		return nil, err
	}
	fmt.Println(user)
	return dto.GetFromCartRequest{UserId: user.Id}, nil
}

func DecodeAddToCartRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.AddToCartRequest

	var user config.JwtCustomClaim

	if err := DecodeJwtHttp(r, "secret", &user); err != nil {
		return req, err
	}
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	req.UserId = user.Id
	return req, nil
}

func DecodeDeleteFromCartRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.DeleteFromCartRequest
	var user config.JwtCustomClaim

	if err := DecodeJwtHttp(r, "secret", &user); err != nil {
		return r, err
	}
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	req.UserId = user.Id
	return req, nil
}
