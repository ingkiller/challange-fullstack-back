package story

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Story struct {
	Id          int
	Title       string
	By          string
	Descendants int
	Kids        []int
	Score       int
	Time        int
	Type        string
	URL         string
}

func GetStoryById(id uint) Story {

	client := &http.Client{}
	storyUrl := fmt.Sprint("https://hacker-news.firebaseio.com/v0/item/", id, ".json")
	req, err := http.NewRequest(http.MethodGet, storyUrl, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var resObject Story
	json.Unmarshal(bodyBytes, &resObject)
	return resObject
}

func GetAll() []Story {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://hacker-news.firebaseio.com/v0/topstories.json", nil)
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

	var responseObject []uint
	json.Unmarshal(bodyBytes, &responseObject)

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
}
