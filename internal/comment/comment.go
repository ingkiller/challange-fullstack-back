package comment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Comment struct {
	Id     int
	Name   string
	Body   string
	PostId int
	Email  string
}

func GetCommentsByPost(postId int) []Comment {
	client := &http.Client{}
	userUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/comments?postId=", postId)
	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var result []Comment
	json.Unmarshal(bodyBytes, &result)
	return result
}

func CountCommentByPost(postId int) int {
	log.Printf("CountCommentByPost: %v", postId)
	return 10
}
