package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type aControl struct {
}

func (con *aControl) Success(c *gin.Context) {

	fmt.Println("success")
}
