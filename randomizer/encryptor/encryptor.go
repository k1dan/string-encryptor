package encryptor

import (
	"context"
	pb "github.com/k1dan/string-encryptor/encryptor/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	address	= "localhost:8080"
)

func Encrypt(strings []string) (encrypted []string) {
	req := &pb.EncryptRequest{Strings: strings}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewEncryptorClient(conn)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	reply, err := client.EncryptString(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	return reply.Strings
}
