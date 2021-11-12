package main

import (
	"fmt"
	"strings"
	"time"
)
//输入一个文本
//打印出文本中某个字母出现的次数，并打印查找操作的时间
//并且将其中的恶臭字符进行打码(114514)
func chazhao(text string) {
	var s string
	fmt.Println("请输入你要查找的字符")
	fmt.Scanln(&s)
	if strings.Count(text, s) != 0 {
		fmt.Println(strings.Count(text, s))
	}else{
		fmt.Println("文本中没有这个字符")
	}
	fmt.Println("进行查找的时间是：",time.Now().Format("2006-01-02 15:04:05"))

}
func main(){
	text:=""
	fmt.Println("请输入一个文本")
	fmt.Scanln(&text)
	ttext:=""
	if strings.Contains(text,"114514") {
		ttext = strings.Replace(text, "114514", "打码", -1)
		fmt.Println("您的文本经过了一定的审核，我们可能会对其中的一些不和谐文字进行一些修改。", "\n", "以下为和谐后的内容", "\n", ttext)
	}else{
		ttext = text
	}
	chazhao(ttext)


}