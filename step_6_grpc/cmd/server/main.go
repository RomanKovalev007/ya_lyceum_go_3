package main

import (
	"example_grpc/include"
	pb "example_grpc/proto"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
  host := "localhost"
  port := "5000"

  addr := fmt.Sprintf("%s:%s", host, port)
  lis, err := net.Listen("tcp", addr)
  if err != nil {
    log.Println("error starting tcp listener: ", err)
    os.Exit(1)
  }
  
  log.Println("tcp listener started at port: ", port)
  grpcServer := grpc.NewServer()
  geomServiceServer := include.NewServer()
  pb.RegisterGeometryServiceServer(grpcServer, geomServiceServer)

  if err := grpcServer.Serve(lis); err != nil {
    log.Println("error serving grpc: ", err)
    os.Exit(1)
  }
}