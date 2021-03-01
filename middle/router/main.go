package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterMiddle1(c *gin.Context)  {
	fmt.Println("我是路由中间件1")
}

func RouterMiddle2(c *gin.Context)  {
	fmt.Println("我是路由中间件2")
}

func oneRouterMiddleHandle() gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println("我是业务处理函数")
	}
}

func main() {
	//创建一个无中间件路由
	router := gin.New()
	// 对于每个路由中间件，您可以添加任意数量的路由中间件
	router.GET("/oneRouterMiddle", RouterMiddle1,RouterMiddle2,oneRouterMiddleHandle())
	// 默认监听本地 0.0.0.0:8080 即localhost:8080 或 127.0.0.1:8080
	router.Run()
}