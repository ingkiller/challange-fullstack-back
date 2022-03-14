package comment

import "log"

type Comment struct {
	Id     uint
	Name   string
	Body   string
	PostId uint
	Email  string
}

func GetCommentsByPost(postId uint) []Comment {
	var result []Comment
	return result
}

func CountCommentByPost(postId int) int {
	log.Printf("CountCommentByPost: %v", postId)
	return 10
}
