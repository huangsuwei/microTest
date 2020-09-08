package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"microProtoTest1/prototext"
)

func main1() {
	text := &prototext.Test{
		Name:    "panda",
		Weight:  []int64{150, 140, 180},
		Stature: 176,
		Motto:   "天行健，君子以自强不息，地势坤，君子以厚德载物",
	}
	fmt.Println(text)
	//proto编码
	data, err := proto.Marshal(text)
	if err != nil {
		fmt.Println("编码失败")
	}
	fmt.Println(data)

	newText := &prototext.Test{}
	err = proto.Unmarshal(data, newText)
	if err != nil {
		fmt.Println("解码失败")
	}
	fmt.Println(newText)
	fmt.Println(newText.String())
	fmt.Println("姓名：", newText.GetName())
	fmt.Println("身高：", newText.GetStature())
	fmt.Println("体重：", newText.GetWeight())
	fmt.Println("格言：", newText.GetMotto())
}
