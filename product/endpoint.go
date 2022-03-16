package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	InsertProductEndpoint  endpoint.Endpoint
	GetAllProductsEndpoint endpoint.Endpoint
	SearchProductEndpoint  endpoint.Endpoint
	PatchProductEndpoint   endpoint.Endpoint
	RemoveProductEndpoint  endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		InsertProductEndpoint:  makeInsertProductEndpoint(s),
		GetAllProductsEndpoint: makeGetAllProductsEndpoint(s),
		SearchProductEndpoint:  makeSearchProductEndpoint(s),
		PatchProductEndpoint:   makePatchProductEndpoint(s),
		RemoveProductEndpoint:  makeRemoveProductEndpoint(s),
	}
}

func makeInsertProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Product)
		err = s.Insert(ctx, req)
		if err != nil {
			return ActionResultProductResponse{Success: false}, err
		}
		return ActionResultProductResponse{Success: true}, nil
	}
}

func makeGetAllProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		resp, serr := s.FindAll(ctx)
		if serr != nil {
			return nil, err
		}
		return resp, nil
	}
}

func makeSearchProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ProductFilterRequest)
		resp, serr := s.Search(ctx, req)
		if serr != nil {
			return nil, err
		}
		return resp, nil
	}
}

func makePatchProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Product)
		err = s.Update(ctx, req)
		if err != nil {
			return ActionResultProductResponse{Success: false}, err
		}
		return ActionResultProductResponse{Success: true}, nil
	}
}

func makeRemoveProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ProductIdRequest)
		err = s.Remove(ctx, req.ProductId)
		if err != nil {
			return ActionResultProductResponse{Success: false}, err
		}
		return ActionResultProductResponse{Success: true}, nil
	}
}
