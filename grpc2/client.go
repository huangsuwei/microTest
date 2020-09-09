package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pd "microTest/myproto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:10086", grpc.WithInsecure())
	if err != nil {
		fmt.Println("客户端网络错误", err)
	}
	//网络延时关闭
	defer conn.Close()

	//获得grpc句柄
	c := pd.NewHelloserverClient(conn)
	//通过句柄调用函数
	re, err1 := c.Sayhello(context.Background(), &pd.HelloReq{Name: "熊猫"})
	if err1 != nil {
		fmt.Println("sayhello调用失败：", err1)
	}
	fmt.Println("sayhello调用结果：", re.Msg)
	re2, err2 := c.Sayname(context.Background(), &pd.NameReq{Name: "托尼帕克"})
	if err2 != nil {
		fmt.Println("Sayname调用失败：", err2)
	}
	fmt.Println("Sayname调用结果：", re2.Msg)
}
