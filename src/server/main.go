package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "gitlab.com/campaignstorage/movingsvc/src/proto"
)

const (
	port = "0.0.0.0:50051"
)

// server is used to implement datamover.FileMover
type FileMoverServer struct{}

func (s *FileMoverServer) Store(ctx context.Context, in *pb.StoreFile) (*pb.MoveStatus, error) {
	return &pb.MoveStatus{Status: 0}, nil
}

func (s *FileMoverServer) Retrieve(ctx context.Context, in *pb.RetrieveFile) (*pb.MoveStatus, error) {
	return &pb.MoveStatus{Status: 1}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFileMoverServer(s, &FileMoverServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("listen on: %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
