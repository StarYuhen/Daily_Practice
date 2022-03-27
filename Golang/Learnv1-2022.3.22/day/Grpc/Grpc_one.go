package main

import (
	pd "./protoc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

// protoc -I . --proto_path=.:$GOPATH/src:../  --go_out=plugins=grpc:. set.proto

// 将proto文件转成pd.go的方式  protoc -I . --go_out=plugins=grpc:. set.proto  必须要进入文件目录才能用这个，否则set.proto 需要加上路径
/*
go_out 指定go语言形式
set.proto指定输入的文件地址
这是同一目录下的生成方式，因为proto指定了输出文件目录，所以才会和当前文件不在一个目录
（set.pd.go 文件）
*/

// ProdenServer 定义约定的接口
type ProdenServer struct{}

// ProdenAdd 添加server服务
var ProdenAdd = ProdenServer{}

const (
	Address = "localhost:2349" // 接口地址
)

// Add 定义函数体方法，Grodens是proto里面的请求结构，Gou则是响应结构
func (Pro ProdenServer) Add(context context.Context, p *pd.Grodens) (*pd.Gou, error) {
	res := new(pd.Gou)
	res.Mess = fmt.Sprintf(p.Name)
	return res, nil
}

func main() {
	fmt.Println("服务开启：")
	// prc就是tcp协议
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		panic(err)
	}

	// 实例化接口
	st := grpc.NewServer()

	// 注册接口
	// pd.RegisterProdenServer(st, ProdenAdd)

	grpclog.Info("list tcp in:", Address)

	err = st.Serve(listen)
	if err != nil {
		grpclog.Error(err)
	}
}
