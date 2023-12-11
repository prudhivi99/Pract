package main

import (
	"context"
	"local/grpc1/greet/proto"
	"log"
)

func doGreet(c proto.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &proto.GreetRequest{
		FirstName: "Praveen",
	})

	if err != nil {
		log.Fatalf("Could  not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
