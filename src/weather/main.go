package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "weather/protos"
	svr "weather/service"
)

func main() {
	lis, err := net.Listen("tcp", ":9020")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &svr.WeatheService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
