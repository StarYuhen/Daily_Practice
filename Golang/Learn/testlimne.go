package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{ServerIp: serverIp, ServerPort: serverPort}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("Net Dial ERROR!:", err)
		return nil
	}
	client.conn = conn
	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "www.baidu.com", "ServerIP:(default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 80, "ServerPort(default:8080)")
}

func main() {
	// 命令行解析
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>Net Connect ERROR!>>>>>>>>>>")
		return
	}
	fmt.Println(">>>>>>Net Connect SUCCESS!>>>>>>>>>>>")
	select {}
}
