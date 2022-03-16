package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type List struct {
	Id        int
	UserId    int
	Title     string
	Completed bool
}

func GetListByUserId(userId int) []List {
	client := &http.Client{}
	userUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/todos?userId=", userId)
	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var result []List
	json.Unmarshal(bodyBytes, &result)
	return result
}

func ToggleTask(taskId int) bool {
	return true
}
