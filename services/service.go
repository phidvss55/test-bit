package services

import (
	"context"
	"encoding/json"
	"go-circleci/types"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*types.CatFact, error)
	
	// Product operations
	GetAllProducts(context.Context) ([]*types.Product, error)
	GetProductByID(ctx context.Context, id int) (*types.Product, error)
	CreateProduct(ctx context.Context, req *types.CreateProductRequest) (*types.Product, error)
	UpdateProduct(ctx context.Context, id int, req *types.UpdateProductRequest) (*types.Product, error)
	DeleteProduct(ctx context.Context, id int) error
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) GetCatFact(ctx context.Context) (*types.CatFact, error) {
	res, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return nil, err
	}

	fact := &types.CatFact{}
	if err := json.NewDecoder(res.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}

// Stub implementations for product methods - will be properly handled by CompositeService in task 9
func (s *CatFactService) GetAllProducts(ctx context.Context) ([]*types.Product, error) {
	return nil, nil
}

func (s *CatFactService) GetProductByID(ctx context.Context, id int) (*types.Product, error) {
	return nil, nil
}

func (s *CatFactService) CreateProduct(ctx context.Context, req *types.CreateProductRequest) (*types.Product, error) {
	return nil, nil
}

func (s *CatFactService) UpdateProduct(ctx context.Context, id int, req *types.UpdateProductRequest) (*types.Product, error) {
	return nil, nil
}

func (s *CatFactService) DeleteProduct(ctx context.Context, id int) error {
	return nil
}
