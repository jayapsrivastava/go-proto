package main

import (
	"log"
	"os"
	"path"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "gitlab.com/campaignstorage/movingsvc/src/proto"

)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFileMoverClient(conn)

	// Contact the server and print out its response.
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %v <source file> <dest file>", path.Base(os.Args[0]))
	}
	srcPath := os.Args[1]
	dstPath := os.Args[2]

	fi, err := os.Stat(srcPath)
	if os.IsNotExist(err) {
		log.Fatal("could not access source path")
	}
	if ! fi.Mode().IsRegular() {
		log.Fatal("source path is not a regular file")
	}

	r, err := c.Store(context.Background(), &pb.StoreFile{SourcePath: srcPath, DestinationPath: dstPath})
	if err != nil {
		log.Fatalf("could not send request: %v", err)
	}
	log.Printf("store status: %v", r.Status)
}
