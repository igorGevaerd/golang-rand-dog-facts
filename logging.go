package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetDogFact(ctx context.Context) (fact *DogFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%v err=%s took=%v", fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetDogFact(ctx)
}
