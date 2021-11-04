package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var td = make(chan int,1)

func add() {
for i := 0; i < 50000; i++ {
x :=<- td
x++
td <- x

}
wg.Done()
}
func main() {
td <- 0
wg.Add(2)
go add()
go add()
wg.Wait()
a :=<- td
fmt.Println(a)
}