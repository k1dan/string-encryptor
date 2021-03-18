package server

import (
	"context"
	"github.com/k1dan/string-encryptor/encryptor/proto"
	"github.com/k1dan/string-encryptor/encryptor/workerpool"
	"os"
	"strconv"
)

func getEnv() (w int) {
	workers := os.Getenv("WORKERS")
	if workers == "" {
		return 4
	}
	wr, _ := strconv.Atoi(workers)
	return wr
}

type GRPCServer struct {}

func (s *GRPCServer) EncryptString(ctx context.Context, req *proto.EncryptRequest) (rep *proto.EncryptReply, err error) {
	workersNumber := getEnv()
	strings := req.Strings
	encrypted := workerpool.PooledWork(strings, workersNumber)

	return &proto.EncryptReply{Encrypted: true, Strings: encrypted}, nil
}