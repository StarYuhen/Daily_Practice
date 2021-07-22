package main

import (
	pb "./protoc" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Addresst Address gRPC服务地址
	Addresst = "127.0.0.1:2349"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Addresst, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			grpclog.Error(err)
		}
	}(conn)

	// 初始化客户端
	c := pb.NewProdenClient(conn)

	// 调用方法
	req := &pb.Grodens{Name: "gRPC"}
	res, err := c.Add(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Info(res.ProtoMessage)
}
