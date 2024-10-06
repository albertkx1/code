package router

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"test01/main/bean"
	"test01/main/controller"
	"test01/main/middleware"
)

func AdminGroup(r *gin.Engine) {
	group := r.Group("/", middleware.Init3)
	{
		group.GET("/", start)
		group.GET("/hello", hello)
		group.POST("/info", getValue)
		group.POST("/xml", xmlInfo)
		group.GET("/user", middleware.Init, controller.UserControl{}.Index)
		group.GET("/load", controller.UserControl{}.Load)
		group.POST("/upload", controller.UserControl{}.Upload)
		group.GET("/loadMore", controller.UserControl{}.LoadMore)
		group.POST("/uploadMore", controller.UserControl{}.UploadMore)
	}
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "%v", "1")
	//cookie, err := c.Cookie("username")
	//if err != nil {
	//	c.String(http.StatusOK, err.Error())
	//}

	c.String(http.StatusOK, fmt.Sprintln(sessions.Default(c).Get("admin")))
	c.JSON(200, gin.H{
		//c.JSON(200, =map[string]interface{}{
		"message": "pong",
	})

}
func xmlInfo(c *gin.Context) {
	user := &bean.Info{}
	data, _ := c.GetRawData()
	if err := xml.Unmarshal(data, user); err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
	}
}
func getValue(c *gin.Context) {
	user := &bean.Info{}
	if err := c.ShouldBind(&user); err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
	}
	//fmt.Println(name + strconv.Itoa(age) + sex + id)

	c.JSON(200, gin.H{
		"user": user,
	})
}
func start(c *gin.Context) {
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"message": "pong",
	//	"msg":     123,
	//})
	c.HTML(http.StatusOK, "web/test01.html", gin.H{
		"Name":  "123",
		"Data":  "456",
		"news":  &bean.A{Name: "hello", Data: "world"},
		"habit": []string{"sleep", "eat", "drink"},
		"date":  1629423555,
	})
	//c.JSON(http.StatusOK, &a{
	//	Name: "123",
	//	Data: "456",
	//})
}
func f_(info [][]int, x, y int) int {
	ans := -1
	for i, v := range info {
		x1 := v[0] + v[2]
		y1 := v[1] + v[3]
		if x <= x1 && x >= v[0] && y >= v[1] && y <= y1 {
			ans = i
		}

	}
	return ans
}

func main() {
	var n, tmp, x, y int
	fmt.Scanf("%v", &n)
	info := make([][]int, n)
	for i := 0; i < n; i++ {
		args := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			fmt.Scanf("%d", &tmp)
			args = append(args, tmp)
		}
		info = append(info, args)
	}
	fmt.Scanf("%v %v", &x, &y)
	fmt.Println(f_(info, x, y))
}
