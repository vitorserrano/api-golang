package db

import "api/golang/pkg"

var commentPostMap = make(map[string]map[string]*pkg.Comment)

type CommentRepository struct {
}

func (CommentRepository) InsertComment(postId string, comment *pkg.Comment) {
	commentMap := commentPostMap[postId]
	if commentMap == nil {
		commentMap = make(map[string]*pkg.Comment, 0)
	}
	commentMap[comment.Id.String()] = comment
	commentPostMap[postId] = commentMap
}

func (CommentRepository) GetComments(postId string) []*pkg.Comment {
	comments := make([]*pkg.Comment, 0)
	for _, v := range commentPostMap[postId] {
		comments = append(comments, v)
	}

	return comments
}

func (CommentRepository) GetCommentById(postId string, commentId string) *pkg.Comment {
	commentMap := commentPostMap[postId]
	return commentMap[commentId]
}
