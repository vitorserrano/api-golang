package http

import (
	"api/golang/internal/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	engine := gin.Default()
	gin.EnableJsonDecoderDisallowUnknownFields()

	engine.GET("/health", handler.HealthCheck)

	posts := engine.Group("/posts")
	posts.GET("/", handler.GetPosts)
	posts.GET("/:id", handler.GetPost)
	posts.POST("/", handler.CreatePost)
	posts.PUT("/:id", handler.UpdatePost)
	posts.PATCH("/:id", handler.PartialUpdatePost)
	posts.DELETE("/:id", handler.DeletePost)

	posts.GET("/:postId/comments", handler.GetComments)
	posts.POST("/:postId/comments", handler.CreateComment)
	posts.GET("/:postId/comments/:id", handler.GetComment)

	return engine.Run(":3000")
}
