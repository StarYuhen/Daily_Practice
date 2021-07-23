package main

import (
	"./protoc"
	"google.golang.org/grpc"
)

type serverFile struct {
	server *grpc.Server
}

func (s *serverFile) Upload(stream protoc.FileUpload_UploadServer) {
}
