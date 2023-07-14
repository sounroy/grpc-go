package main

import (
	"context"
	"log"
	"time"

	pb "com.ea/greet"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Long Greet client was invoked")

	reqs := []*pb.GreetRequest{

		{FirstName: " Sourav"},
		{FirstName: " Roy"},
		{FirstName: " Shahrukh"},
		{FirstName: " Khan"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet L %v \n", err)
	}

	for _, req := range reqs {
		log.Printf(" Sending req : %v \n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error happened while receiving response from LongGreet : %v \n", err)
	}

	log.Printf(" Long Greeet %s \n", res.Result)
}
