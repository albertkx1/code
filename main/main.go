package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"test01/main/router"
	"time"
)

func timestamp(t int) string {
	fmt.Println(t)
	unix := time.Unix(int64(t), 0)
	return unix.Format("2006-01-02 15:04:05")
}
func stringCat(t1 string, t2 string) string {
	fmt.Println("sc")
	return t1 + t2
}

func main() {
	fmt.Println("123")
	r := gin.Default()
	r.Static("/staticFiles", "./staticFiles")
	//创建存储引擎 密钥
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	//r.Use(middleware.Init, middleware.Init2)
	r.SetFuncMap(template.FuncMap{
		"UnixTime":  timestamp,
		"stringCat": stringCat,
	})
	r.LoadHTMLGlob("template/**/*")
	router.AdminGroup(r)
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
