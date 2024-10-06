package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Init(c *gin.Context) {
	fmt.Println("init")
	start := time.Now().Unix()
	c.Next()
	end := time.Now().Unix()
	if end-start > 0 {
		c.Abort()
	}
	fmt.Println(end - start)
}
func Init2(c *gin.Context) {
	fmt.Println("init2")
	start := time.Now().Unix()
	c.Next()
	end := time.Now().Unix()

	fmt.Println(end - start)
}
func Init3(c *gin.Context) {
	fmt.Println("init33")
	//fmt.Println(c.Request.URL)
	c.String(200, c.Request.URL.String()+"\n")
	c.Set("admin", 123)
	//context := c.Copy()
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("goroutine done" + context.Request.URL.Path)
	//}()
}
