package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go-circleci/repository"
	"go-circleci/types"
)

// ProductService implements the Service interface for product operations
type ProductService struct {
	repo repository.ProductRepository
}

// NewProductService creates a new ProductService with the given repository
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetAllProducts retrieves all products from the repository
func (s *ProductService) GetAllProducts(ctx context.Context) ([]*types.Product, error) {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}
	
	// Return empty slice instead of nil for consistency
	if products == nil {
		return []*types.Product{}, nil
	}
	
	return products, nil
}

// GetProductByID retrieves a single product by its ID with validation
func (s *ProductService) GetProductByID(ctx context.Context, id int) (*types.Product, error) {
	// Validate ID
	if id <= 0 {
		return nil, errors.New("invalid product ID: must be greater than 0")
	}
	
	product, err := s.repo.GetByID(ctx, id)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("product with ID %d not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	
	return product, nil
}

// CreateProduct creates a new product with input validation
func (s *ProductService) CreateProduct(ctx context.Context, req *types.CreateProductRequest) (*types.Product, error) {
	// Validate required fields
	if req.Name == "" {
		return nil, errors.New("product name is required")
	}
	
	if req.Price < 0 {
		return nil, errors.New("product price must be greater than or equal to 0")
	}
	
	if req.Stock < 0 {
		return nil, errors.New("product stock must be greater than or equal to 0")
	}
	
	// Create product entity
	product := &types.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	
	// Call repository to create
	if err := s.repo.Create(ctx, product); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}
	
	return product, nil
}

// UpdateProduct updates an existing product with validation
func (s *ProductService) UpdateProduct(ctx context.Context, id int, req *types.UpdateProductRequest) (*types.Product, error) {
	// Validate ID
	if id <= 0 {
		return nil, errors.New("invalid product ID: must be greater than 0")
	}
	
	// Validate required fields
	if req.Name == "" {
		return nil, errors.New("product name is required")
	}
	
	if req.Price < 0 {
		return nil, errors.New("product price must be greater than or equal to 0")
	}
	
	if req.Stock < 0 {
		return nil, errors.New("product stock must be greater than or equal to 0")
	}
	
	// Create product entity with ID
	product := &types.Product{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	
	// Call repository to update
	if err := s.repo.Update(ctx, product); err == sql.ErrNoRows {
		return nil, fmt.Errorf("product with ID %d not found", id)
	} else if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}
	
	return product, nil
}

// DeleteProduct deletes a product by its ID with validation
func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	// Validate ID
	if id <= 0 {
		return errors.New("invalid product ID: must be greater than 0")
	}
	
	// Call repository to delete
	if err := s.repo.Delete(ctx, id); err == sql.ErrNoRows {
		return fmt.Errorf("product with ID %d not found", id)
	} else if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	
	return nil
}

// GetCatFact is a stub implementation to satisfy the Service interface
// This will be properly handled by CompositeService in task 9
func (s *ProductService) GetCatFact(ctx context.Context) (*types.CatFact, error) {
	return nil, errors.New("GetCatFact not supported by ProductService")
}
