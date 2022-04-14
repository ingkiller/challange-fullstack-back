package comment

import (
	"encoding/json"
	"fmt"
	"github.com/ingkiller/hackernews/internal/client"
	"io/ioutil"
)

type Comment struct {
	Id     int
	Name   string
	Body   string
	PostId int
	Email  string
}

var CommentsByPostId = make(map[int][]Comment)

func GetCommentsByPost(postId int) []Comment {

	if CommentsByPostId[postId] != nil {
		return CommentsByPostId[postId]
	}

	commentUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/comments?postId=", postId)
	resp, err := client.MakeReq(commentUrl)
	if err != nil {
		fmt.Print("Error: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error: %v", err.Error())
	}
	var result []Comment
	json.Unmarshal(bodyBytes, &result)
	CommentsByPostId[postId] = result
	return result
}
