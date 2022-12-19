package main

import (
	"context"
	pb "grpc/proto"
	"io"
	"log"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names : %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()
	for _, name := range names.Names {
		res := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
