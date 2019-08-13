package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":  "testHandlers",
		"name":     "clarkrao",
		"age":      "24",
		"nickName": "CK",
	})
}

func TestUrlParam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func TestGet(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "Guest")
	lastName := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

	c.JSON(http.StatusOK, gin.H{
		"firstName": firstName,
		"lastName":  lastName,
	})
}

func TestPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

	c.JSON(200, gin.H{
		"method":  http.MethodPost,
		"status":  http.StatusOK,
		"message": message,
		"nick":    nick,
	})
}

func TestReflect(c *gin.Context) {
	val := 3.1
	v := reflect.ValueOf(val)

	c.JSON(200, gin.H{
		"type":    v.Type().String(),
		"kind":    v.Kind(),
		"value":   v.Float(),
		"isFloat": v.Kind() == reflect.Float64,
	})
}
