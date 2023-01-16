package main

import (
	pd "./protoc"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// PdServer 注册微服务结构体
type PdServer struct {
	pd.UnimplementedTestServer
}

func (s *PdServer) TestHello(ctx context.Context, request *pd.TestHello) *pd.TestReturn {
	return &pd.TestReturn{
		Return: "Hello" + request.Name,
	}
}

// grpc使用
func main() {
	// 使用gin绑定微服务接口
	engine := gin.New()
	// 启动pd服务
	PdEngine := grpc.NewServer()
	pd.RegisterTestServer(PdEngine, &PdServer{})
	engine.Run(":678")
}
