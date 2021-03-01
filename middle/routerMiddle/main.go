package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GroupRouterGoodsMiddle1(c *gin.Context)  {
	fmt.Println("我是goods路由组中间件1")
}

func GroupRouterGoodsMiddle2(c *gin.Context) {
	fmt.Println("我是goods路由组中间件2")
}

func GroupRouterOrderMiddle1(c *gin.Context) {
	fmt.Println("我是order路由组中间件1")
}

func GroupRouterOrderMiddle2(c *gin.Context) {
	fmt.Println("我是order路由组中间件2")
}

func main() {
	//创建一个无中间件路由
	router := gin.New()
	router.Use(gin.Logger())

	//第1种路由组使用方式 可以添加多个处理函数 但是不知道为什么 官方举例的这第一种方式用不了
	router.Group("/goods", GroupRouterGoodsMiddle1, GroupRouterGoodsMiddle2)
	router.GET("/goods/add", func(context *gin.Context) {
		fmt.Println("/goods/add")
	})


	//第2种路由组使用方式
	orderGroup := router.Group("/order")
	orderGroup.Use(GroupRouterOrderMiddle1, GroupRouterOrderMiddle2)
	{
		orderGroup.GET("/add", func(context *gin.Context) {
			fmt.Println("/order/add")

		})

		orderGroup.GET("/del", func(context *gin.Context) {
			fmt.Println("/order/del")

		})

		//orderGroup下再嵌套一个testGroup
		testGroup:=orderGroup.Group("/test", func(context *gin.Context) {
			fmt.Println("order/test下的中间件")
		})

		testGroup.GET("/test1", func(context *gin.Context) {
			fmt.Println("order/test/test1的函数")
		})
	}

	router.Run()
}
