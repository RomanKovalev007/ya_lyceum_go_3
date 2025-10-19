package include

import (
	"context"
	pb "example_grpc/proto"
	"log"
)

type Server struct {
  pb.GeometryServiceServer // сервис из сгенерированного пакета
}

func NewServer() *Server {
  return &Server{}
}

type GeometryServiceServer interface {
  Area(context.Context, *pb.RectRequest) (*pb.AreaResponse, error)
  Perimeter(context.Context, *pb.RectRequest) (*pb.PermiterResponse, error)
}

func (s *Server) Area(
  ctx context.Context, 
  in *pb.RectRequest,
) (*pb.AreaResponse, error) {
  log.Println("invoked Area: ", in)
  return &pb.AreaResponse{
    Result: in.Height * in.Width,
  }, nil
}

func (s *Server) Perimeter(
  ctx context.Context, 
  in *pb.RectRequest,
) (*pb.PermiterResponse, error) {
  log.Println("invoked Perimeter: ", in)
  return &pb.PermiterResponse{
    Result: 2 * (in.Height + in.Width),
  }, nil
}