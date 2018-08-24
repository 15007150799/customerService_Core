package main

import (
	"git.jsjit.cn/customerService/customerService_Core/controller"
	_ "git.jsjit.cn/customerService/customerService_Core/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	r := gin.Default()

	defaultController := controller.InitHealth()
	dialogController := controller.InitDialog()
	serverController := controller.InitServer()
	offlineReplyController := controller.InitOfflineReply()

	// 定义路由
	v1 := r.Group("/v1")
	{
		// 健康检查
		v1.GET("/health", defaultController.Health)
		v1.GET("/init", defaultController.Init)

		// 会话操作
		dialog := v1.Group("/dialog")
		{
			dialog.GET("list", dialogController.List)
			dialog.POST("create", dialogController.Create)
			dialog.GET("customer/:id/history", dialogController.History)
			dialog.POST("customer/:id/message", dialogController.SendMessage)
			dialog.DELETE("customer/:id/message", dialogController.RecallMessage)
		}

		// 客服操作
		server := v1.Group("/server")
		{
			server.GET(":id", serverController.Get)
			server.POST(":id/status", serverController.ChangeStatus)
		}

		// 设置操作
		setting := v1.Group("/setting")
		{
			// 离线自动回复设置
			offlineReply := setting.Group("offline_reply")
			{
				offlineReply.GET("", offlineReplyController.List)
				offlineReply.POST("", offlineReplyController.Create)
				offlineReply.PUT(":id", offlineReplyController.Update)
				offlineReply.DELETE(":id", offlineReplyController.Delete)
			}
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":5000")
}
