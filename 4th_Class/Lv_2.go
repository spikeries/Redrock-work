package main

import (
	"fmt"
	"sync"
	"time"
)
var xiecheng sync.WaitGroup
func timeR1() {

	t1 := ""
	for {
		t1 = time.Now().Format(time.Kitchen)
		if t1 == "2:00AM" {
			fmt.Println("谁能比我卷？！")
		}
	}
}
func timeR2(){
	t1 := ""
	for {
		t1 = time.Now().Format(time.Kitchen)
		if t1 == "6:00AM" {
			fmt.Println("早八算什么，早六才是吾辈应起之时")
		}
	}
}
func timeR3() {
	for {
		fmt.Println("芜湖！起飞！")
		time.Sleep(30*time.Second)
	}
	}

func main(){
	xiecheng.Add(3)
	go timeR1()
	go timeR2()
	go timeR3()
	xiecheng.Wait()
}
