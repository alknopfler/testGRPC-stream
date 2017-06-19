package main

import (
	"log"
	pb "github.com/alknopfler/testGRPC-stream/requester"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"bufio"
	"os"
	"fmt"
	"io"
)

const (
	address     = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRequesterClient(conn)

	var key string
	for key != "default"{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter key: ")
		key, _ := reader.ReadString('\n')

		r, err := c.Process(context.Background(), &pb.Request{KeyId: key})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		for {
			flow, err := r.Recv()
			if err == io.EOF {
				fmt.Println("entra y rompe")
				break

			}

			log.Printf("Output from server: %s", flow)
		}

	}

}

