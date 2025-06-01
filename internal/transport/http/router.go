package http

import (
	_ "github.com/adamknbrv/forum/docs"
	"github.com/adamknbrv/forum/internal/transport/http/handlers"
	"github.com/adamknbrv/forum/pkg/ws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// @title Forum API
// @version 1.0
// @description API для форума с обсуждениями и сообщениями
// @host localhost:8082
func SetupRouters(handler handlers.Handler, hub *ws.Hub) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	go hub.Run()

	messageRoutes := router.Group("/messages")
	{
		messageRoutes.GET("/:id", handler.GetAllMessage)
		messageRoutes.GET("/user/:id", handler.GetMessageByUserID)
		messageRoutes.POST("", handler.CreateMessage)
		messageRoutes.PUT("", handler.UpdateMessage)
		messageRoutes.DELETE("/:id", handler.DeleteMessage)
	}

	discussionRoutes := router.Group("/discussions")
	{

		discussionRoutes.GET("", handler.GetAllDiscussion)
		discussionRoutes.GET("/:id", handler.GetDiscussionByID)
		discussionRoutes.GET("/user/:id", handler.GetDiscussionsByUserID)
		discussionRoutes.POST("", handler.CreateDiscussions)
		discussionRoutes.PUT("/update/:id", handler.UpdateDiscussions)
		discussionRoutes.DELETE("/:id", handler.DeleteDiscussions)
		discussionRoutes.GET("/chat/:id", hub.Chat)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
