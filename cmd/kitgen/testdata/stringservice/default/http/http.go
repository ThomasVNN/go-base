package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ThomasVNN/go-base/cmd/kitgen/testdata/stringservice/default/endpoints"
	gotransport "github.com/ThomasVNN/go-base/transport/http"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/concat", gotransport.NewServer(endpoints.Concat, DecodeConcatRequest, EncodeConcatResponse))
	m.Handle("/count", gotransport.NewServer(endpoints.Count, DecodeCountRequest, EncodeCountResponse))
	return m
}
func DecodeConcatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ConcatRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeConcatResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeCountResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
