package main

import (
  "context"
  "log"
  "time"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  pb "main/proto"
)

func main() {
  conn, err := grpc.NewClient(
    "localhost:50051",
    grpc.WithTransportCredentials(insecure.NewCredentials()),
  )
  if err != nil {
    log.Fatalf("dial error: %v", err)
  }
  defer conn.Close()

  client := pb.NewGreeterClient(conn)
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()

  resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
  if err != nil {
    log.Fatalf("RPC error: %v", err)
  }
  log.Printf("Greeting: %s", resp.GetMessage())
}