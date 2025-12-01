package services

import (
	"context"
	"go-circleci/types"
)

// CompositeService wraps both CatFactService and ProductService to implement the full Service interface
type CompositeService struct {
	catFactService *CatFactService
	productService *ProductService
}

// NewCompositeService creates a new CompositeService with both CatFact and Product capabilities
func NewCompositeService(catFactService *CatFactService, productService *ProductService) Service {
	return &CompositeService{
		catFactService: catFactService,
		productService: productService,
	}
}

// GetCatFact delegates to the CatFactService
func (s *CompositeService) GetCatFact(ctx context.Context) (*types.CatFact, error) {
	return s.catFactService.GetCatFact(ctx)
}

// GetAllProducts delegates to the ProductService
func (s *CompositeService) GetAllProducts(ctx context.Context) ([]*types.Product, error) {
	return s.productService.GetAllProducts(ctx)
}

// GetProductByID delegates to the ProductService
func (s *CompositeService) GetProductByID(ctx context.Context, id int) (*types.Product, error) {
	return s.productService.GetProductByID(ctx, id)
}

// CreateProduct delegates to the ProductService
func (s *CompositeService) CreateProduct(ctx context.Context, req *types.CreateProductRequest) (*types.Product, error) {
	return s.productService.CreateProduct(ctx, req)
}

// UpdateProduct delegates to the ProductService
func (s *CompositeService) UpdateProduct(ctx context.Context, id int, req *types.UpdateProductRequest) (*types.Product, error) {
	return s.productService.UpdateProduct(ctx, id, req)
}

// DeleteProduct delegates to the ProductService
func (s *CompositeService) DeleteProduct(ctx context.Context, id int) error {
	return s.productService.DeleteProduct(ctx, id)
}
