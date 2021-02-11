package main

import (
	"log"
	"net"

	"github.com/MudimukkuSreenath/go-grpc-Sreenath/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Failed to listen on port 9000: %v", err)
	}
	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to sere grpc server over port 9000: %v", err)

	}
}
