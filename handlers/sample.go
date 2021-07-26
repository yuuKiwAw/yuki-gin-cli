package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "HelloWorld",
	})
}

func SampleMW(c *gin.Context) {
	req, _ := c.Get("request")
	fmt.Println("request22", req)
	c.JSON(200, gin.H{
		"request": req,
	})
}
