package server

import (
	"context"
	"github.com/k1dan/string-encryptor/encryptor/proto"
	"github.com/k1dan/string-encryptor/encryptor/workerpool"
)

type GRPCServer struct {}

func (s *GRPCServer) EncryptString(ctx context.Context, req *proto.EncryptRequest) (rep *proto.EncryptReply, err error) {
	strings := req.Strings
	encrypted := workerpool.PooledWork(strings, 4)

	return &proto.EncryptReply{Encrypted: true, Strings: encrypted}, nil
}