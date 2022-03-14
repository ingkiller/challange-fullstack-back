package post

import (
	"encoding/json"
	"fmt"
	"github.com/ingkiller/hackernews/internal/user"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	Id     int
	Title  string
	Body   string
	UserId int
	User   user.User
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
	log.Printf("Unmarshaled: %v", responseObject)
	var result []Post
	for i := 0; i < 3; i++ {
		var temp user.User
		temp = user.GetUserById(responseObject[i].UserId)
		newPost := responseObject[i]
		newPost.User = temp
		result = append(result, newPost)

	}
	log.Printf("result: %v", result)
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
