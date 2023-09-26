package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"net/http"
	"web-template/endpoints"

	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler")
)

func NewHTTPHandler(endpoints endpoints.Endpoints, logger log.Logger) http.Handler {
	m := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerErrorLogger(logger),
	}

	m.Methods("GET").Path("/health").Handler(httptransport.NewServer(endpoints.HealthEndpoint, DecodeHttpHealthRequest, EncodeHttpGenericResponse, options...))
	m.Methods("GET").Path("/greeting").Handler(httptransport.NewServer(endpoints.GreetingEndpoint, DecodeHttpGreetingRequest, EncodeHttpGenericResponse, options...))

}

func EncodeHttpGenericResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoints.Failer); ok && f.Failed() != nil {
		encodeError(ctx, f.Failed(), writer)
		return nil
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(writer).Encode(response)
}

func DecodeHttpHealthRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return endpoints.HealthRequest{}, nil
}

func DecodeHttpGreetingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := r.URL.Query()
	names, exists := vars["name"]
	if !exists || len(names) != 1 {
		return nil, ErrBadRouting
	}
	req := endpoints.GreetingRequest{Name: names[0]}
	return req, nil
}

type errorWrapper struct {
	Error string `json:"error"`
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func err2code(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
