package main

import (
	"fmt"
	"local/grpc1/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(in *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was inviled with: %v\n", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&proto.GreetResponse{
			Result: res,
		})
	}
	return nil
}
