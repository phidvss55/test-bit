package logger

import (
	"context"
	"fmt"
	"go-circleci/services"
	"go-circleci/types"
	"time"
)

type LoggingService struct {
	next services.Service
}

func NewLoggingService(next services.Service) services.Service {
	return &LoggingService{next: next}
}

func (s *LoggingService) GetCatFact(context context.Context) (fact *types.CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%v took=%v", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(context)
}

func (s *LoggingService) GetAllProducts(ctx context.Context) (products []*types.Product, err error) {
	defer func(start time.Time) {
		fmt.Printf("GetAllProducts count=%d err=%v took=%v\n", len(products), err, time.Since(start))
	}(time.Now())

	return s.next.GetAllProducts(ctx)
}

func (s *LoggingService) GetProductByID(ctx context.Context, id int) (product *types.Product, err error) {
	defer func(start time.Time) {
		fmt.Printf("GetProductByID id=%d err=%v took=%v\n", id, err, time.Since(start))
	}(time.Now())

	return s.next.GetProductByID(ctx, id)
}

func (s *LoggingService) CreateProduct(ctx context.Context, req *types.CreateProductRequest) (product *types.Product, err error) {
	defer func(start time.Time) {
		fmt.Printf("CreateProduct name=%s err=%v took=%v\n", req.Name, err, time.Since(start))
	}(time.Now())

	return s.next.CreateProduct(ctx, req)
}

func (s *LoggingService) UpdateProduct(ctx context.Context, id int, req *types.UpdateProductRequest) (product *types.Product, err error) {
	defer func(start time.Time) {
		fmt.Printf("UpdateProduct id=%d name=%s err=%v took=%v\n", id, req.Name, err, time.Since(start))
	}(time.Now())

	return s.next.UpdateProduct(ctx, id, req)
}

func (s *LoggingService) DeleteProduct(ctx context.Context, id int) (err error) {
	defer func(start time.Time) {
		fmt.Printf("DeleteProduct id=%d err=%v took=%v\n", id, err, time.Since(start))
	}(time.Now())

	return s.next.DeleteProduct(ctx, id)
}
