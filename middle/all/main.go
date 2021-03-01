package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GlobalMiddle(c *gin.Context){
	fmt.Println("我是全局中间件")
}

func main() {
	//创建一个无中间件路由
	router := gin.New()

	//使用自定义的全局中间件
	router.Use(GlobalMiddle)

	router.GET("/", func(context *gin.Context) {
		fmt.Println("我是/")
	})

	router.Run()
}
