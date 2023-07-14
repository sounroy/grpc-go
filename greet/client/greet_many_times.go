package main

import (
	"context"
	"io"
	"log"

	pb "com.ea/greet"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println(" Many times do greet client was invoked")

	req := &pb.GreetRequest{
		FirstName: "Sourav",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf(" Error while calling the greet many time server : %v \n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf(" Error while reading the stream %v \n", err)
		}

		log.Printf(" Greet many times %s \n", msg.Result)

	}
}
