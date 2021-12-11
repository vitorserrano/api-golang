package handler

import (
	"api/golang/internal/container"
	"api/golang/pkg"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user pkg.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, pkg.NewResponseError("Invalid user", err))
		return
	}

	container.UserRepository.CreateUser(&user)

	c.JSON(201, pkg.ResponseUser{
		Login: user.Login,
	})
}
