package main

import (
	"fmt"
	"log"

	pb "com.ea/greet"
)

// Copy this from greet_grpc.pb.go
// GreetmanyTimes(*GreetRequest, GreetService_GreetmanyTimesServer) error
// Add the receiver function to denote that this function belongs to server & add pb to request response
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf(" Greeting many times function was callwd with %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf(" Hello Bro %s, mumber %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
