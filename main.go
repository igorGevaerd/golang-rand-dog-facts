package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := NewDogFactService("https://dogapi.dog/api/v2/facts?limit=1")
	svc = NewLoggingService(svc)

	fact, err := svc.GetDogFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", fact)
	fmt.Println("xisd")
}
