package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); !ok {
		fmt.Println(e, ok)
		w.WriteHeader(403)
		return e
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}
