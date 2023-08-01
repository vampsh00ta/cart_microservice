package http

import (
	"cart_mircoservice/iternal/service/dto"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"

	//"github.com/gofrs/uuid/v5"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func DecodeGetFromCart(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	userIdStr, ok := vars["uuid"]
	if !ok {
		return nil, ErrBadRouting
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return nil, err
	}
	return dto.GetFromCartRequest{UserId: userId}, nil
}

func DecodeAddToCartRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.AddToCartRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func DecodeDeleteFromCartRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req dto.DeleteFromCartRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}
