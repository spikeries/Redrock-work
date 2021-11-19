package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)
var users = make(map[string]string,0)
var check = make(map[string]string,0)
func main(){
	r:=gin.Default()
	f1,_:=os.OpenFile("users.data",os.O_CREATE,0666)
	r.POST("/login",func (c *gin.Context){
		d,_:=ioutil.ReadFile("users.data")
		json.Unmarshal(d,&check)
		uname:=c.PostForm("username")
		pword:=c.PostForm("password")
	if check[uname]==pword||pword!=""{
		c.SetCookie("login_cookie",uname,3600,"/","",false,true)
		c.String(200,"登录成功")
	}else{
		c.String(403,"登录失败")
	}
	})
	r.POST("/register",func (c *gin.Context){
		d,_:=ioutil.ReadFile("users.data")
		json.Unmarshal(d,&check)
		uname:=c.PostForm("username")
		pword:=c.PostForm("password")
		_,che:=check[uname]
		if che {
			c.String(403,"用户名已被注册")
			c.Abort()
		}
		if len(pword)<6{
			c.String(403,"密码不得小于六位")
		}
		users[uname]=pword
		userdata,_:=json.Marshal(users)
		f1.Write(userdata)
c.String(200,"注册成功，请重新登录。")
	})
	check:=func(c *gin.Context){

			value,err1 := c.Cookie("login_cookie")
			if err1!=nil {
				c.String(403, "认证失败，妹有cookie")
				c.Abort()
			
		}
		c.Set("cookie",value)
	}
	r.GET("/hello",check,func(c *gin.Context){
		cookie,_:=c.Get("cookie")
		s:=cookie.(string)

			c.String(200, "登录成功,你好"+s)

	})
	r.Run()
	defer f1.Close()
}
