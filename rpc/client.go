package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("网络连接失败！", err)
	}

	var pd int
	//func (p *Panda) GetInfo(argType int, replyType *int) error {
	err = cli.Call("Panda.GetInfo", 10086, &pd)
	if err != nil {
		fmt.Println("方法调用失败！")
	}
	fmt.Println("最后得到的值：", pd)
}
