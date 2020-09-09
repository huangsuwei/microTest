package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	pd "microTest/myproto"
	"net"
)

type server struct{}

func (this *server) Sayhello(ctx context.Context, in *pd.HelloReq) (out *pd.HelloRsp, err error) {
	return &pd.HelloRsp{Msg: "hello" + in.Name}, err
}

func (this *server) Sayname(ctx context.Context, in *pd.NameReq) (out *pd.NameRsp, err error) {
	return &pd.NameRsp{Msg: in.Name + "早上好！"}, err
}

func main() {
	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误", err)
	}
	//创建grpc的服务
	srv := grpc.NewServer()
	pd.RegisterHelloserverServer(srv, &server{})

	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络有问题", err)
	}
}
