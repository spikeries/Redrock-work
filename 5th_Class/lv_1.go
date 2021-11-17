package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
r:=gin.Default()
check:=func(c *gin.Context){
value,err:=c.Cookie("login_cookie")
if err!=nil{
	c.String(403,"认证失败，妹有cookie")
	c.Abort()
}
c.Set("cookie",value)
	}
r.POST("/login",func (c *gin.Context){
	uname:=c.PostForm("username")
	pword:=c.PostForm("password")
	if uname == "wanghao" && pword == "whnb"{
		c.SetCookie("login_cookie",uname,3600,"/","",false,true)
		c.String(200,"登录成功")
	}else{
		c.String(403,"登录失败")
	}
})
r.GET("/hello",check,func(c *gin.Context){
	cookie,_:=c.Get("cookie")
	s:=cookie.(string)
c.String(200,"登录成功,你好"+s)
})
r.Run()
}