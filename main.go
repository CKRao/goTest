package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goTestProject/handler"
	"net/http"
)

const PORT = ":6789"

func main() {
	logrus.Info("Start Server")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//初始化路由
	initRouter(r)

	r.Run(PORT) // listen and serve on 127.0.0.1:6789
}

//初始化路由
func initRouter(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/file"
		r.HandleContext(c)
	})

	r.GET("/test", handler.Test)

	r.GET("/testReflect", handler.TestReflect)

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	r.GET("/user/:name", handler.TestUrlParam)

	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", handler.TestGet)

	//测试POST请求
	r.POST("/form_post", handler.TestPost)

	//静态资源路径
	r.Static("/assets", "/opt")

	//文件系统
	r.StaticFS("/file", http.Dir("/opt/download"))

	//图标
	r.StaticFile("/favicon.ico", "./favicon.ico")
}
