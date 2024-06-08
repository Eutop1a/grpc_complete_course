package main

import (
	"flag"
	"fmt"
	"gitlab.com/techschool/pcbook/pb"
	"gitlab.com/techschool/pcbook/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Printf("start server port %d", *port)

	laptopServer := service.NewLaptopSever(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf(":%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("cannot start listen: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start serve: %v", err)
	}
}
