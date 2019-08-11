package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"goTestProject/handler"
)

func main() {
	logrus.Info("Start Server")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", handler.Test)

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	r.GET("/user/:name", handler.TestUrlParam)

	// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", handler.TestGet)

	//测试POST请求
	r.POST("/form_post", handler.TestPost)

	r.Run(":6789") // listen and serve on 127.0.0.1:6789
}
