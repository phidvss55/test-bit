package services

import (
	"context"
	"encoding/json"
	"go-circleci/types"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*types.CatFact, error)
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
