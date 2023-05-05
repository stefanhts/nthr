package main

import (
    pb "nthr/server/pb/cabinet"
    "log"
    "net"

    "google.golang.org/grpc"
)

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
    return &pb.GetResponse{Value: "Hello " + in.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterCabinetServer(s, &server{})
    s.Serve(lis)
}
