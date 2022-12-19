package main

import (
	"context"
	pb "grpc/proto"
	"log"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client stream started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Not able to send names %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		//time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished ")
	if err != nil {
		log.Fatalf("Error while reciving %v ", err)
	}
	log.Printf("%v", res.Messages)
}
