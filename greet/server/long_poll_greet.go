package main

import (
	"fmt"
	"io"
	"log"

	pb "com.ea/greet"
)

func (s *Server) longGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println(" The Long greet function was invoked by client")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf((""))
		}
		
		res += fmt.Sprint(" Hello %s ! \n", req.FirstName)

}}