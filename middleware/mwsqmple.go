package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func SampleMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件测试")
		c.Set("request", "中间件测试")
		c.Next()
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
