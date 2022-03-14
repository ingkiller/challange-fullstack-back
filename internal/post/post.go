package post

import (
	"encoding/json"
	"fmt"
	"github.com/ingkiller/hackernews/internal/comment"
	"github.com/ingkiller/hackernews/internal/user"
	"io/ioutil"
	"net/http"
	"sync"
)

type Post struct {
	Id              int
	Title           string
	Body            string
	UserId          int
	User            user.User
	NumberOfComment int
}

func GetAll() []Post {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print("read all error: %v", err.Error())
	}

	var responseObject []Post
	json.Unmarshal(bodyBytes, &responseObject)
	//	var posts []Post
	//log.Printf("Unmarshaled: %v", responseObject)
	var result []Post
	ch := make(chan user.User)
	chComment := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(3)
		for j := 0; j < 3; j++ {
			go func(p Post) {
				defer wg.Done()
				ch <- user.GetUserById(p.UserId)

			}(responseObject[j])
		}
		wg.Wait()
		close(ch)
	}()

	go func() {
		var wgComment sync.WaitGroup
		wgComment.Add(3)
		for j := 0; j < 3; j++ {
			go func(p Post) {
				defer wgComment.Done()
				chComment <- comment.CountCommentByPost(p.Id)
			}(responseObject[j])
		}
		wgComment.Wait()
		close(chComment)
	}()

	var users []user.User
	for c := range ch {
		users = append(users, c)
	}

	var numberOfComment []int
	for c := range chComment {
		numberOfComment = append(numberOfComment, c)
	}

	for i := 0; i < 3; i++ {
		newPost := responseObject[i]
		newPost.User = users[i]
		newPost.NumberOfComment = numberOfComment[i]
		result = append(result, newPost)
	}
	//	log.Printf("result: %v", result)
	return result
	/*
		var stories []Story
		c := make(chan Story)
		go func() {
			var wg sync.WaitGroup
			wg.Add(3)
			for j := 0; j < 3; j++ {
				go func(s uint) {
					defer wg.Done()
					c <- GetStoryById(s)
				}(responseObject[j])
			}
			wg.Wait()
			close(c)
		}()
		for v := range c {
			stories = append(stories, v)
		}
		log.Printf("Unmarshaled: %v", stories)
		return stories
	*/
}
