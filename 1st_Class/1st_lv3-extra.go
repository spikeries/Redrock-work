package main

import (
	"fmt"
	"math/rand"
)
func bulb(s []int)[]int{
	length := len(s)
	var i int
	for i = 0 ; i <= length - 1 ; i++{
		for j := 0 ; j <= length - 2 ; j++{
		if s[j] < s[j+1]{
			s[j],s[j+1] = s[j+1],s[j]
		}

		}

	}
return s
}

func main(){
var s1 []int
s1 = make([]int,0)
for i := 0; i <= 99; i++{
var a = rand.Intn(1000)
s1 = append(s1,a)
}
s1 = bulb (s1)

fmt.Print(s1)
}
