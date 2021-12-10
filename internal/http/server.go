package http

import (
	"api/golang/internal/handler"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	engine := gin.Default()
	gin.EnableJsonDecoderDisallowUnknownFields()

	engine.GET("/health", handler.HealthCheck)
	engine.GET("/posts", handler.GetPosts)
	engine.POST("/posts", handler.CreatePost)
	engine.PUT("/posts/:id", handler.UpdatePost)
	engine.PATCH("/posts/:id", handler.PartialUpdatePost)

	return engine.Run(":3000")
}
