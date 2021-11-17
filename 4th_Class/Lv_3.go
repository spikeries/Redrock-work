package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var users = make(map[string]string,0)
var check = make(map[string]string,0)
func reg(){
	f1,err1:=os.OpenFile("User.data",os.O_CREATE,0666)
	d,_:=ioutil.ReadFile("User.data")
	json.Unmarshal(d,&check)
	if err1 != nil{
		return}
	defer f1.Close()
for {
	fmt.Println("注册：请输入用户名")
	reader1 := bufio.NewReader(os.Stdin)
	uname, _ := reader1.ReadString('\n')
	_, che := check[uname]
	if che {
		fmt.Println("该用户名已存在，请重新输入")
		continue
	}

for {
	fmt.Println("注册：请输入密码（密码不得少于6位）")
	reader2 := bufio.NewReader(os.Stdin)
	upassword, _ := reader2.ReadString('\n')
	if len(upassword) <= 6 {
		fmt.Println("密码不得少于6位！请重新输入密码。")
continue
	}
	users[uname] = upassword
	userdata, err2 := json.Marshal(users)
	if err2 != nil {
		fmt.Println("注册失败")
		return
	}

	_, err3 := f1.Write(userdata)
	if err3 != nil {
		fmt.Println("写入失败")
		return
	}
	break
}
	break
}
}
func log(){
	f1,err1:=os.OpenFile("User.data",os.O_CREATE,0666)
	d,_:=ioutil.ReadFile("User.data")
	json.Unmarshal(d,&check)
	if err1 != nil{
		return}
	defer f1.Close()
for {
	fmt.Println("登录：请输入用户名")
	reader1 := bufio.NewReader(os.Stdin)
	uname, _ := reader1.ReadString('\n')
	_, che := check[uname]
	if !che {
		fmt.Println("用户名不存在，请重新输入")
		continue
	}
for {
	fmt.Println("登录：请输入密码")
	reader2 := bufio.NewReader(os.Stdin)
	upassword, _ := reader2.ReadString('\n')
	if upassword == check[uname] {
		fmt.Println("登录成功~~~")
	} else {
		fmt.Println("登录失败...", "\n", "请再次输入密码")
		continue
	}
	break
}
	break
}

}
func main(){
	loop:
	for {
		fmt.Println("注册请按1，登录请按2，退出请按3")
		var a int
		fmt.Scanln(&a)
		switch a {
		case 1:
			{
				reg()
			}
		case 2:
			{
				log()
			}
		case 3:
			{
			break loop
			}
		default:
			{
				fmt.Println("没有这个选项哦")
			}
			continue
		}
	}


}
