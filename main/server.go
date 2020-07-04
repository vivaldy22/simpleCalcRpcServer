package main

import (
	"github.com/edwardsuwirya/simpleCalcRpcServer/api"
	"github.com/edwardsuwirya/simpleCalcRpcServer/calc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := calc.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterAdditionServer(grpcServer, &s)
	log.Println("Server runs")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
