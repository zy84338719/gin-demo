package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个无中间件路由
	router := gin.New()
	router.Run()
}
