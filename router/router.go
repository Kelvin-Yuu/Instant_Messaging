package router

import (
	"im/middlewares"
	"im/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//用户登录
	r.POST("/login", service.Login)

	auth := r.Group("/u", middlewares.AuthCheck())

	//用户详情
	auth.GET("/user/detail", service.UserDetail)

	return r
}
