package main

import (
	"fmt"
	"sync"
)
var wa sync.WaitGroup
var a = make(chan int,1)
func add1() {
	var x int
	x=<-a
	x++
	fmt.Println("这是协程1：", x)
	a<-x
}
func add2(){
	var x int
	//time.Sleep(1*time.Millisecond)
	x= <- a
	x++
	fmt.Println("这是协程2:",x)
    a<-x
	wa.Done()
}

func main() {
	a <- 0
	for i := 0; i < 50; i++ {
		wa.Add(1)
		add1()
		go add2()
		wa.Wait()
	}
}


