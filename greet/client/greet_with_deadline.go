package main

import (
	"context"
	"log"
	"time"

	pb "com.ea/greet"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadLine(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println(" doGreetWithDeadLine was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{

		FirstName: "Sourav",
	}

	res, err := c.GreetWithDeadLine(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded !")
				return
			} else {
				log.Fatalf("Unxepected grpc error : %v \n", err)
			}
		} else {
			log.Fatalf(" A non gRPC error happened %v \n", err)
		}

	}
	log.Printf(" Greet with deadline : %s \n", res.Result)
}
