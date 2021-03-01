package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//自定义中间件第1种定义方式
func CustomRouterMiddle1(c *gin.Context)  {
	t := time.Now()
	fmt.Println("我是自定义中间件第1种定义方式---请求之前")
	//在gin上下文中定义一个变量
	c.Set("example", "CustomRouterMiddle1")
	//请求之前
	c.Next()
	fmt.Println("我是自定义中间件第1种定义方式---请求之后")
	//请求之后
	//计算整个请求过程耗时
	t2 := time.Since(t)
	log.Println(t2)

}

//自定义中间件第2种定义方式
func CustomRouterMiddle2() gin.HandlerFunc{
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("我是自定义中间件第2种定义方式---请求之前")
		//在gin上下文中定义一个变量
		c.Set("example", "CustomRouterMiddle2")
		//请求之前
		c.Next()
		fmt.Println("我是自定义中间件第2种定义方式---请求之后")
		//请求之后
		//计算整个请求过程耗时
		t2 := time.Since(t)
		log.Println(t2)
	}
}



func main() {

	r := gin.New()

	//测试时下面两个中间件选择一个，注释一个
	r.Use(CustomRouterMiddle1)
	r.Use(CustomRouterMiddle2())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		//
		log.Println(example)
	})

	// 监听本地8080端口
	r.Run(":8080")
}
