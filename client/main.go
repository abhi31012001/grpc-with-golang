package main

import (
	"fmt"
	pb "grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8088"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Not able to connect to grpc server %v", err)
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	names := &pb.NamesList{
		Names: []string{"Abhi", "Traves", "Virat"},
	}

	//		callSayHello(client)
	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	v := time.Now().Local()
	for i := 0; i < 100000; i++ {
		callSayHelloBidirectionalStream(client, names)
	}
	fmt.Print(v)
	fmt.Print(time.Now().Local())
}
