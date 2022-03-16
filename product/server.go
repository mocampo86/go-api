package product

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/product").Handler(httpTransport.NewServer(
		endpoints.GetAllProductsEndpoint,
		DecodeGetRequest,
		EncodeResponse,
	))

	r.Methods("GET").Path("/product/search").Handler(httpTransport.NewServer(
		endpoints.SearchProductEndpoint,
		DecodeProductFilterRequest,
		EncodeResponse,
	))

	r.Methods("POST").Path("/product").Handler(httpTransport.NewServer(
		endpoints.InsertProductEndpoint,
		DecodePostProductRequest,
		EncodeResponse,
	))

	r.Methods("PATCH").Path("/product").Handler(httpTransport.NewServer(
		endpoints.PatchProductEndpoint,
		DecodePostProductRequest,
		EncodeResponse,
	))

	r.Methods("DELETE").Path("/product/{id}").Handler(httpTransport.NewServer(
		endpoints.RemoveProductEndpoint,
		DecodeProductIdRequest,
		EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
