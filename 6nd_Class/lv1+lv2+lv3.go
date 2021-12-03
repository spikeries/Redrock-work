package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

var db, err = sql.Open("mysql", "root:123456@/user")
var buffer bytes.Buffer
type user struct {
	id       int
	age      int
	name     string
	username string
	password string
	question string
	answer   string
	email    string
}
type comment struct {
	id       int
	reid     int
	sender   string
	receiver string
	message string
}
var message = make([]string,0)
func main() {
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("/board",board)
	r.POST("/login", login)
	r.POST("/regist", regist)
	r.POST("/find", find)
	r.POST("/update", update)
	r.POST("/message",me)
	check1 := func(c *gin.Context) {

		value, err1 := c.Cookie("login_cookie")
		if err1 != nil {
			c.String(403, "认证失败，没有cookie")
			c.Abort()

		}
		c.Set("cookie", value)
	}
	r.GET("/hello", check1, func(c *gin.Context) {
		cookie, _ := c.Get("cookie")
		s := cookie.(string)

		c.String(200, "登录成功,你好"+s)

	})
	r.Run()
}

func login(c *gin.Context) {
	_, err := c.Cookie("login_cookie")
	if err == nil {
		c.String(403, "您已登录，请勿重复登录。")
		c.Abort()
		return
	}
	uname := c.PostForm("username")
	pword := c.PostForm("password")
	che := "select username,password from users where id > ?"
	stmt, err := db.Prepare(che)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if uname == u.username && pword == u.password && uname != "" {
			c.SetCookie("login_cookie", uname, 1800, "/", "", false, true)
			c.String(http.StatusOK, "登录成功")
			return
		}
	}
		c.String(403, "用户名或密码错误，请重新登录。")

		return


}

func regist(c *gin.Context) {
	_, err := c.Cookie("login_cookie")
	if err == nil {
		c.String(403, "您已登录，无法进行注册。")
		return
	}
	name := c.PostForm("name")
	age := c.PostForm("age")
	uname := c.PostForm("username")
	pword := c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	email := c.PostForm("email")

	if len(uname) < 6 {
		c.String(403, "用户名不得小于6位，请重新注册。")
		c.Abort()
		return
	}
	if question == "" || answer == "" {
		c.String(403, "请设置密保")
		c.Abort()
		return
	}
	//在这里限定了各项参数，防止和数据库中的参数限制冲突。
	if uname == "" || pword == "" || name == "" || age == "" || email == "" {
		c.String(403, "value不能为空，请重新注册。")
		c.Abort()
		return
	}
	if check(db, uname, email) {
		c.String(403, "出现错误，用户名或邮箱已被注册，请重新注册。")
		c.Abort()
		return
	}
	sqlStr := "insert into users(name,age,username,password,question,answer,email) values (?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, age, uname, pword, question, answer, email)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	c.String(403, "注册成功，请重新登录。")
	return
}

func check(db *sql.DB, reguser, email string) bool {
	check := "select username,email from users where id > ?"
	c, err := db.Prepare(check)
	if err != nil {
		fmt.Println("prepare failed")
		return true
	}
	defer c.Close()
	rows, err := c.Query(0)
	if err != nil {
		fmt.Println("query failed")
		return true
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.username, &u.email)
		if err != nil {
			fmt.Println("scan filed")
			return true
		}
		if reguser == u.username || email == u.email {
			return true
		}
	}
	return false
}

func find(c *gin.Context) {
	q := c.PostForm("question")
	a := c.PostForm("answer")
	e := c.PostForm("email")
	if q == "" || a == "" || e == "" {
		c.String(403, "请输入密保,注册邮箱,找回账号")
	}
	sqlStr := "select username,question,answer,email from users where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.username, &u.question, &u.answer, &u.email)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if u.question == q && u.answer == a {
			c.String(200, "密保与邮箱输入正确，您的账号是"+u.username)
			return
		}
	}
}

func update(c *gin.Context) {
	un := c.PostForm("username")
	q := c.PostForm("question")
	a := c.PostForm("answer")
	p := c.PostForm("newpassword")
	if p == "" {
		c.String(403, "请输入新密码,newpassword")
	}
	if q == "" || a == "" || un == "" {
		c.String(403, "请输入密保,用户名,修改密码")
	}
	sqlStr := "select username,question,answer from users where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.username, &u.question, &u.answer)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if u.question == q && u.answer == a && u.username == un {
			sqlStr := "UPDATE users SET password=? WHERE username=?"
			stmt, err := db.Prepare(sqlStr)
			if err != nil {
				fmt.Printf("prepare sql failed, err:%v\n", err)
				return
			}
			_, err = stmt.Exec(p, un)
			if err != nil {
				fmt.Printf("exec failed, err:%v\n", err)
				return
			}
			c.String(200, "您的密码已修改完成。")
			return
		}
	}
}

func board(c *gin.Context) {
	receiver, err := c.Cookie("login_cookie")
	if err != nil {
		c.String(403, "请先登录再查看留言与回复哦。")
		return
	}
	sqlStr := "select reid,id,sender,receiver,message from comment where receiver = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
		rows, err := stmt.Query(receiver)
	    m := ""
		for rows.Next() {
			var co comment
			err := rows.Scan(&co.reid, &co.id, &co.sender, &co.receiver, &co.message)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
			a := strconv.Itoa(co.id)
			m += "来自"+co.sender+"的信息:"+"\n"+co.message+"\n"+"本条信息的id为: "+a+"\n"
			if co.reid!=0{
				a := strconv.Itoa(co.reid)
				m += "回复的信息的id为"+"\n"+a+"\n"
			}
		}
		if m == ""{
			c.String(200,"留言板：目前还没有人留言或回复哦~")
		}
		c.String(200,"留言板：\n%s",m)
		}

func me(c *gin.Context){
	receiver:=c.PostForm("receiver")
	//如果是留言需要写，回复不需要，
	// 会根据回复的信息的id自动读取出receiver
	reid:=c.PostForm("reid")
	//如果是留言则不需要填写reid
	//如果是回复则需要在reid中写出要回复的信息的id
	message:=c.PostForm("message")
	if message==""{
		c.String(403,"留言内容不能为空哦~")
	}
	sender,err:=c.Cookie("login_cookie")
	if err != nil{
		c.String(403,"请先进行登录。")
	}
	if reid != "" {
		//reid不等于0，判定为对某条信息的回复。
		reid, err := strconv.Atoi(reid)
		if err != nil {
			c.String(403, "要回复的id非法。")
		}
		sqlStr := "select id,sender from comment where id > ?"
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			fmt.Printf("prepare failed, err:%v\n", err)
			return
		}
		defer stmt.Close()
		rows, err := stmt.Query(0)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return
		}
		defer rows.Close()
		// 循环读取结果集中的数据
		for rows.Next() {
			var co comment
			err := rows.Scan(&co.id,&co.receiver)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
            if co.id == reid {
				sqlStr = "insert into comment(receiver,sender,reid,message) values (?,?,?,?)"
				stmt, err = db.Prepare(sqlStr)
				if err != nil {
					fmt.Println("prepare failed, err:%v\n", err)
				}
				defer stmt.Close()
				_, err = stmt.Exec(co.receiver, sender, reid, message)
				c.String(200,"回复成功！")
				break
			}
			c.String(403,"似乎并不存在你要回复的id号哦~")
		}
	}
	//如果reid不填，则判定为留言。
	//需要填出接收者
	if reid == ""{
		//检查留言的接收者是否存在。
		sqlStr := "select username from users where id > ?"
		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			fmt.Printf("prepare failed, err:%v\n", err)
			return
		}
		defer stmt.Close()
		rows, err := stmt.Query(0)
		if err != nil {
			fmt.Printf("query failed, err:%v\n", err)
			return
		}
		defer rows.Close()
		// 循环读取结果集中的数据
		for rows.Next(){
			var u user
			err := rows.Scan(&u.username)
			if err != nil {
				fmt.Printf("scan failed, err:%v\n", err)
				return
			}
			if u.username == receiver{
				sqlStr = "insert into comment(sender,receiver,message) values (?,?,?)"
				stmt, err = db.Prepare(sqlStr)
				if err != nil {
					fmt.Printf("prepare failed, err:%v\n", err)
					return
				}
				defer stmt.Close()
				_, err = stmt.Exec(sender,receiver,message)
				if err != nil {
					fmt.Printf("insert failed, err:%v\n", err)
					return
				}
				c.String(200,"留言成功~")
				return
			}
			}
		c.String(403,"你要进行留言的用户不存在。")
		return
	}
}