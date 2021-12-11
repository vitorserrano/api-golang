package db

import "api/golang/pkg"

var userMap = make(map[string]*pkg.User, 0)

type UserRepository struct{}

func (UserRepository) CreateUser(user *pkg.User) {
	userMap[user.Login] = user
}

func (UserRepository) GetUser(login string) *pkg.User {
	return userMap[login]
}
