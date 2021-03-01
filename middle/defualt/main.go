package main

import "github.com/gin-gonic/gin"

func main() {
	// 默认启动方式，包含 Logger、Recovery 中间件
	router:=gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"name":"zhangyi",
		})
	})
	router.Run()
}
