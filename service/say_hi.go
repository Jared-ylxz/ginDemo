package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sayhi(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": fmt.Sprintf("hello, world. my name is %s", name),
	})
}