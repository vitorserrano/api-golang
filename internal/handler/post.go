package handler

import (
	"api/golang/pkg"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var postMap = make(map[string]*pkg.Post, 0)

func GetPosts(c *gin.Context) {
	titleParam := c.Query("title")

	posts := make([]pkg.Post, 0)
	for _, v := range postMap {
		if len(titleParam) > 0 && !strings.Contains(v.Title, titleParam) {
			continue
		}

		posts = append(posts, *v)
	}

	c.JSON(200, posts)
}

func GetPost(c *gin.Context) {
	post, responseError := findPost(c)

	if responseError != nil {
		c.JSON(404, responseError)
		return
	}

	c.JSON(200, post)
}

func CreatePost(c *gin.Context) {
	requestPost, responseError := parseBody(c)
	if responseError != nil {
		c.JSON(400, responseError)
		return
	}

	post := pkg.Post{
		Id:       uuid.New(),
		Title:    requestPost.Title,
		Body:     requestPost.Body,
		User:     requestPost.User,
		DateTime: time.Now(),
	}

	postMap[post.Id.String()] = &post
	log.Println(fmt.Sprintf("post %s created", post))

	c.JSON(201, post)
}

func UpdatePost(c *gin.Context) {
	post, responseError := findPost(c)
	if responseError != nil {
		c.JSON(404, responseError)
		return
	}

	requestPost, responseError := parseBody(c)
	if responseError != nil {
		c.JSON(400, responseError)
		return
	}

	post.Title = requestPost.Title
	post.Body = requestPost.Body
	post.User = requestPost.User
	post.DateTime = time.Now()

	c.JSON(200, post)
}

func PartialUpdatePost(c *gin.Context) {
	post, responseError := findPost(c)
	if responseError != nil {
		c.JSON(404, responseError)
		return
	}

	var partialRequest pkg.PartialRequestPost
	if err := c.ShouldBindJSON(&partialRequest); err != nil {
		if err == io.EOF {
			log.Println("[WARN] -  empty json", err)
			c.JSON(400, pkg.NewResponseError("empty json", err))
			return
		}

		log.Println("[WARN] -  invalid json", err)
		c.JSON(400, pkg.NewResponseError("invalid json", err))
		return
	}

	if partialRequest.Title != nil {
		post.Title = *partialRequest.Title
	}

	if partialRequest.Body != nil {
		post.Body = *partialRequest.Body
	}

	if partialRequest.User != nil {
		post.User = *partialRequest.User
	}

	c.JSON(204, "")
}

func DeletePost(c *gin.Context) {
	post, responseError := findPost(c)

	if responseError != nil {
		c.JSON(404, responseError)
		return
	}

	delete(postMap, post.Id.String())
	c.JSON(204, "")
}

func parseBody(c *gin.Context) (*pkg.RequestPost, *pkg.ResponseError) {
	var requestPost pkg.RequestPost

	if err := c.ShouldBindJSON(&requestPost); err != nil {
		if err == io.EOF {
			log.Println("[WARN] -  empty json", err)
			return nil, pkg.NewResponseError("empty json", err)
		}

		log.Println("[WARN] -  invalid json", err)
		return nil, pkg.NewResponseError("invalid json", err)
	}

	return &requestPost, nil
}

func findPost(c *gin.Context) (*pkg.Post, *pkg.ResponseError) {
	id := c.Param("id")
	post := postMap[id]
	if post == nil {
		return nil, &pkg.ResponseError{
			Cause:   "id not found",
			Message: fmt.Sprintf("id %s not found", id),
		}
	}

	return post, nil
}
