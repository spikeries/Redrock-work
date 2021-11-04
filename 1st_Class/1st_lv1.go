package main

import "fmt"

func main(){
 var str string
 fmt.Scanln(&str)
	 ob := []byte(str)
x := 0
x = len(ob)
i := 0
for i = 0; i <= (x-1)/2; i++{
	ob[i],ob[x - i - 1]=ob[x - i - 1],ob[i]
}


	 os := string(ob)//转成什么括号外就写什么

	 fmt.Print (os)

}
