package main

import (
	"log"
	"net"

	pb "github.com/alknopfler/testGRPC-stream/requester"
	"google.golang.org/grpc"
	"time"
)

const (
	port = ":50051"
)

type server struct {}

func (s *server) Process( in *pb.Request,stream pb.Requester_ProcessServer) error {
	m := make(map[string]string)
	m["key1"]="value1"
	m["key2"]="value2"
	m["key3"]="value3"
	m["key4"]="value4"
	m["default"] = "not found...Quit"
	for i := range m{
			if err := stream.Send(&pb.Response{Val: &pb.Value{Value: m[i]}}); err != nil{
				return err
			}
		time.Sleep(1*time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRequesterServer(s, &server{})
	s.Serve(lis)
}

