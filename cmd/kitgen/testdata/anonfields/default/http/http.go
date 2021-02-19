package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ThomasVNN/go-base/cmd/kitgen/testdata/anonfields/default/endpoints"
	gotransport "github.com/ThomasVNN/go-base/transport/http"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/foo", gotransport.NewServer(endpoints.Foo, DecodeFooRequest, EncodeFooResponse))
	return m
}
func DecodeFooRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.FooRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeFooResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
