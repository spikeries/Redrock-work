package main

import "fmt"

type studata struct {
	name string
	number int
}
type stulist []studata
type num struct {
	x int
}
type numlist []num


var a = stulist{
	{
		name: "小明",
		number: 3,
	},
	{
		name:"小华",
		number: 1,
	},
	{
		name:"小强",
		number: 2,
	},

}
var b = numlist{
	{x:3},
	{x:1},
	{x:2},
}

func (o numlist)sort(){
	for i:=0;i<len(o);i++{
		for j:=0;j<len(o)-1;j++{
			if o[j].x<o[j+1].x{
				o[j],o[j+1]=o[j+1],o[j]
			}
		}
	}

}
func (o stulist)sort(){
	for i:=0;i<len(o);i++{
		for j:=0;j<len(o)-1;j++{
			if o[j].number<o[j+1].number{
				o[j],o[j+1]=o[j+1],o[j]
			}
		}
	}

}
type p interface {
	sort()
}
func ssort(s p){
s.sort()
}
func main(){
	ssort(a)
	ssort(b)
	fmt.Println(a)
	fmt.Println(b)

}