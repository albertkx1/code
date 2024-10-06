package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type UserControl struct {
	aControl
}

func (con UserControl) Index(c *gin.Context) {
	time.Sleep(1 * time.Second)
	value, _ := c.Get("admin")
	session := sessions.Default(c)
	session.Set("admin", 123456)
	if session.Save() != nil {
		fmt.Println("error saving session")
	}
	//c.SetCookie("alt", "6", 3600, "/", "localhost", false, true)
	//fmt.Println(c.Cookie("alt"))
	sprint := fmt.Sprintln(value)
	c.String(200, "hello"+sprint)
	con.Success(c)
}
func (con UserControl) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	ext := filepath.Ext(file.Filename)
	m := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := m[ext]; !ok {
		c.String(200, "not support")
	}
	day := time.Now().Format("20060102")
	fmt.Println(day)
	dir := "./file/" + day
	fmt.Println(dir)
	if err := os.MkdirAll(dir, 0666); err != nil {
		fmt.Println(err)
		c.String(200, "mkdir error")
	}
	err := c.SaveUploadedFile(file, dir+"/"+file.Filename)
	if err != nil {
		c.String(400, err.Error())
	} else {
		c.String(200, "ok")
	}
}
func (con UserControl) Load(c *gin.Context) {
	c.HTML(200, "web/upload.html", gin.H{})
}

func (con UserControl) UploadMore(c *gin.Context) {
	files, _ := c.MultipartForm()
	//log.Println(files.Filename)
	headers := files.File["file[]"]
	i := 0
	for _, file := range headers {

		i++
		log.Println(file.Filename)
		err := c.SaveUploadedFile(file, "./file/"+strconv.Itoa(i)+file.Filename)
		if err != nil {
			c.String(400, err.Error())
		} else {
			c.String(200, "ok")
		}
	}

}
func (con UserControl) LoadMore(c *gin.Context) {
	c.HTML(200, "web/uploadmore.html", gin.H{})

}
