package main

import (
	"fmt"
    "math/rand"
 "sort"
)

func main(){
 var s1 []int
 s1 = make([]int,0)
 for i := 0; i <= 99; i++{
  var a = rand.Intn(1000)
  s1 = append(s1,a)

 }
 sort.Sort(sort.IntSlice(s1))
fmt.Print(s1)
}
