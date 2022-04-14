package post

import (
	"encoding/json"
	"fmt"
	"github.com/ingkiller/hackernews/graph/model"
	"github.com/ingkiller/hackernews/internal/client"
	"github.com/ingkiller/hackernews/internal/comment"
	"github.com/ingkiller/hackernews/internal/user"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Post struct {
	Id              int
	Title           string
	Body            string
	UserId          int
	User            user.User
	CreatedDate     time.Time
	NumberOfComment int
}

var PostArr []Post

func addInfoToPost(responseObject []Post) []Post {
	var result []Post
	ch := make(chan user.User)
	chComment := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(responseObject))
		for j := 0; j < len(responseObject); j++ {
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
		wgComment.Add(len(responseObject))
		for j := 0; j < len(responseObject); j++ {
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
	for i := 0; i < len(responseObject); i++ {
		newPost := responseObject[i]
		newPost.User = users[i]
		newPost.NumberOfComment = numberOfComment[i]
		randomTime := time.Now().Unix() - rand.Int63n(4406400)
		newPost.CreatedDate = time.Unix(randomTime, 0)
		result = append(result, newPost)
	}
	PostArr = result
	return result
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

	var result []Post
	ch := make(chan user.User)
	chComment := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(responseObject))
		for j := 0; j < len(responseObject); j++ {
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
		wgComment.Add(len(responseObject))
		for j := 0; j < len(responseObject); j++ {
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
	for i := 0; i < len(responseObject); i++ {
		newPost := responseObject[i]
		newPost.User = users[i]
		newPost.NumberOfComment = numberOfComment[i]
		randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
		newPost.CreatedDate = time.Unix(randomTime, 0)
		result = append(result, newPost)
	}
	PostArr = result
	return result
}

func GetPostByRange(start int, long int) []Post {
	var result []Post

	fmt.Print("len(PostArr)r: %v", len(PostArr))
	if len(PostArr) > 0 {
		if (start + long) <= len(PostArr) {
			result = PostArr[start:long]
			fmt.Print("result: %v", PostArr[start:long])
		}
	} else {
		var temp = GetAll()
		fmt.Print("temp: %v", len(temp))
		if len(temp) > 0 {
			if (start + long) <= len(temp) {
				result = PostArr[start:long]
			}
		}
	}
	return result
}

func GetPostsByUserId(userId int, start int, long int) []*model.Post {
	postUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/posts?userId=", userId)
	if userId == 0 {
		postUrl = fmt.Sprint("https://jsonplaceholder.typicode.com/posts")
	}
	resp, err := client.MakeReq(postUrl)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("Error: %v", err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var resObject []Post
	json.Unmarshal(bodyBytes, &resObject)
	resObject = addInfoToPost(resObject)
	var result []*model.Post
	for _, post := range resObject {
		result = append(result, &model.Post{
			ID:              post.Id,
			Title:           post.Title,
			Body:            post.Body,
			NumberOfComment: post.NumberOfComment,
			CreatedDate:     post.CreatedDate.String(),
			User:            &model.User{Name: post.User.Name, Username: post.User.Username, Website: post.User.Website, Email: post.User.Email},
		})
	}
	return result[start:long]
}
