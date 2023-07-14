package main

import (
	"log"
	"time"

	pb "com.ea/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50011"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(" Failed to connect : %v \n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	doGreetManyTimes(c)
	doLongGreet(c)
	doGreetWithDeadLine(c, 5*time.Second)
}
