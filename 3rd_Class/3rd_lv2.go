package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	cj, err := os.Create("plan.txt")
	if err != nil {
		fmt.Println("创建文件失败")
		return
	}
	defer cj.Close()
	_, er := cj.WriteString("I’m not afraid of difficulties and insist on learning programming")
	if er != nil {
		fmt.Println("写入失败")
		return
	}
	red,err1:=os.Open("plan.txt")
	if err1 != nil{
		fmt.Println("打开失败")
		return
	}
	defer red.Close()
	duqu := make([]byte,1,1)
	for {
		n, err2 := red.Read(duqu)//每次读取一个byte
		if n==0||err2==io.EOF{
			fmt.Println("\n读取完啦！")
			break
		}
		fmt.Print(string(duqu[:n]))
	}
}