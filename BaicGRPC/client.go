package main

import (
	"log"

	"github.com/MudimukkuSreenath/go-grpc-Sreenath/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Heelo from the client!",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling sayhello: %s", err)
	}
	log.Printf("response from server: %s", response.Body)
}
