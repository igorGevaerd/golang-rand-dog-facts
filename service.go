package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetDogFact(context.Context) (*DogFact, error)
}

type DogFactService struct {
	url string
}

func NewDogFactService(url string) Service {
	return &DogFactService{
		url: url,
	}
}

func (s *DogFactService) GetDogFact(ctx context.Context) (*DogFact, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &DogData{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	fact := &DogFact{}
	fact.Fact = data.Data[0].Attributes.Fact

	return fact, nil
}
