package main

import (
	"log"
)

func main() {
	svc := NewDogFactService("https://dogapi.dog/api/v2/facts?limit=1")
	svc = NewLoggingService(svc)

	// fact, err := svc.GetDogFact(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", fact)

	ApiServer := NewApiServer(svc)
	log.Fatal(ApiServer.Start(":3000"))

}
