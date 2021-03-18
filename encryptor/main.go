package main

import (
	"github.com/k1dan/string-encryptor/encryptor/proto"
	"github.com/k1dan/string-encryptor/encryptor/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}
	proto.RegisterEncryptorServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
