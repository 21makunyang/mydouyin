package main

import (
	"github.com/gin-gonic/gin"
	"mydouyin/controller"
)

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("mydouyin")

	//basic apis
	apiRouter.POST("/user/register", controller.Register)
	apiRouter.POST("/user/login", controller.Login)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)
}
