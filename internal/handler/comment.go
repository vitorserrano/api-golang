package handler

import (
	"api/golang/internal/container"
	"api/golang/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {

}

func GetComment(c *gin.Context) {

}

func CreateComment(c *gin.Context) {
	postId := c.Param("postId")

	if post := container.PostRepository.FindById(postId); post == nil {
		c.JSON(404, &pkg.ResponseError{
			Cause:   "Post id not found",
			Message: fmt.Sprintf("id &s not found", postId),
		})
	}
}
