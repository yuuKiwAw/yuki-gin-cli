package main

import (
	"yuki-gin-cli/handlers"
	"yuki-gin-cli/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	sample := r.Group("sample")
	{
		sample.GET("/helloworld", handlers.HelloWorld)
		sample.GET("/SampleMW", middleware.SampleMiddleWare(), handlers.SampleMW)
	}

	r.Run(":8080")
}
