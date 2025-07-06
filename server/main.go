package main

import (
  "context"
  "log"
  "net"

  "google.golang.org/grpc"
  pb "main/proto"
)

type greeterServer struct {
  pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(
  ctx context.Context, req *pb.HelloRequest,
) (*pb.HelloReply, error) {
  return &pb.HelloReply{
    Message: "Hello, " + req.GetName(),
  }, nil
}

func main() {
  lis, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatalf("listen error: %v", err)
  }
  srv := grpc.NewServer()
  pb.RegisterGreeterServer(srv, &greeterServer{})
  log.Println("server listening on :50051")
  if err := srv.Serve(lis); err != nil {
    log.Fatalf("serve error: %v", err)
  }
}