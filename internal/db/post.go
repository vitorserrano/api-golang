package db

import (
	"api/golang/pkg"
	"strings"
)

var postMap = make(map[string]*pkg.Post, 0)

type PostRepository struct{}

func (PostRepository) FindPosts(title string) []*pkg.Post {
	posts := make([]*pkg.Post, 0)

	for _, v := range postMap {
		if len(title) > 0 && !strings.Contains(v.Title, title) {
			continue
		}

		posts = append(posts, v)
	}

	return posts
}

func (PostRepository) InsertPost(post *pkg.Post) {
	postMap[post.Id.String()] = post
}

func (PostRepository) UpdatePost(post *pkg.Post) {
	// ToDO
}

func (PostRepository) PartialUpdate(post *pkg.Post) {
	// TODO
}

func (PostRepository) DeletePost(id string) {
	delete(postMap, id)
}

func (PostRepository) FindById(id string) *pkg.Post {
	return postMap[id]
}
