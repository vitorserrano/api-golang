package http

import (
	"api/golang/internal/handler"
	"api/golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	engine := gin.Default()

	gin.EnableJsonDecoderDisallowUnknownFields()
	engine.Use(middleware.TokenValidate())

	engine.GET("/health", handler.HealthCheck)
	engine.POST("/users", handler.CreateUser)
	engine.POST("/login", handler.Login)

	posts := engine.Group("/posts")
	posts.GET("/", handler.GetPosts)
	posts.POST("/", handler.CreatePost)
	posts.GET("/:id", handler.GetPost)
	posts.PUT("/:id", handler.UpdatePost)
	posts.PATCH("/:id", handler.PartialUpdatePost)
	posts.DELETE("/:id", handler.DeletePost)
	posts.POST("/:id/comments", handler.CreateComment)
	posts.GET("/:id/comments", handler.GetCommentsByPost)
	posts.GET("/:id/comments/:commentId", handler.GetCommentByCommentAndPost)

	return engine.Run(":3000")
}
