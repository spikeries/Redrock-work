package main

import (
	"fmt"
	"sync"
)
var xiecheng sync.WaitGroup
func main() {
	//over := make(chan bool) 没有缓存，ture输入造成阻塞。
	over := make(chan bool,1)//添加缓存
	for i := 0; i < 10; i++ {
		//go func() {
		//	fmt.Println(i)
		//}()
		//在主协程退出后自动停止了该子协程，导致无法打印
		//所以要使用sync.WaitGroup
		xiecheng.Add(1)
		go func() {
			fmt.Println(i)
			xiecheng.Done()//结束子协程
		}()
		if i == 9 {
			over <- true
		}
		xiecheng.Wait()//等待打印结束
	}
	<-over
	fmt.Println("over!!!")
}