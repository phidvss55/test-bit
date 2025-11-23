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
