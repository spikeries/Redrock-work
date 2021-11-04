package main

import (
	"fmt"
)

//局部变量会覆盖全局变量
func jia(inputNum1,inputNum2 int) int {
	a := inputNum1 + inputNum2
	return a
}
func jian(inputNum1,inputNum2 int) int {
	return inputNum1 - inputNum2
}
func cheng(inputNum1,inputNum2 int) int {
	return inputNum1 * inputNum2
}
func chu(inputNum1,inputNum2 int) int {
	return inputNum1 / inputNum2
}



func main(){
	var s1 []int
	s1 = make([]int,0)
	l1:
	for {
	inputNum1,inputNum2 := 0,0
	cmd := ""

	fmt.Scanln(&inputNum1, &cmd, &inputNum2)
	var ans int

	switch cmd {
	case "+": ans = jia (inputNum1, inputNum2)
	case "-": ans = jian (inputNum1, inputNum2)
	case "*": ans = cheng (inputNum1, inputNum2)
	case "/": ans = chu (inputNum1, inputNum2)
	default : fmt.Println("字都能打歪来，你褪裙罢")
	goto l1




	}
	fmt.Println(ans)

	s1 = append (s1, ans)
	fmt.Println(s1," ")
	}


}
