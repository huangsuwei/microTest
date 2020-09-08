package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

type Panda int

func (p *Panda) GetInfo(argType int, replyType *int) error {
	fmt.Println("打印对端发送过来的内容为：", argType)
	//修改内容值
	*replyType = argType + 12306

	return nil
}

func main() {
	http.HandleFunc("/panda", pandatext)

	//将类实例化为对象
	pd := new(Panda)
	//服务端注册
	rpc.Register(pd)
	rpc.HandleHTTP()

	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("连接失败！")
	}
	//下面自带高并发
	http.Serve(ln, nil)
}

func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world! hello panda")
}
