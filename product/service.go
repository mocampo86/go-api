package product

import (
	"context"
)

//Service provides some date capabilities to our application
type Service interface {
	Remove(ctx context.Context, id int64) error
	Search(ctx context.Context, filter ProductFilterRequest) (products []Product, err error)
	FindAll(ctx context.Context) (products []Product, err error)
	Update(ctx context.Context, product Product) error
	Insert(ctx context.Context, product Product) error
}

type service struct {
	repository Repository
}

//NewService makes a new Service.
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) Remove(ctx context.Context, id int64) error {
	return s.repository.Remove(ctx, id)
}

func (s service) Search(ctx context.Context, filter ProductFilterRequest) (products []Product, err error) {
	resp, err := s.repository.Search(ctx, filter)
	if err != nil {
		return nil, err
	}
	result := []Product{}
	for _, product := range resp {
		result = append(result, product.ToModel())
	}

	return result, nil
}

func (s service) FindAll(ctx context.Context) (products []Product, err error) {
	resp, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	result := []Product{}
	for _, product := range resp {
		result = append(result, product.ToModel())
	}

	return result, nil
}

func (s service) Update(ctx context.Context, product Product) error {
	return s.repository.Update(ctx, product.ToDAO())
}

func (s service) Insert(ctx context.Context, product Product) error {
	return s.repository.Insert(ctx, product.ToDAO())
}
