package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Id       int
	Name     string
	Username string
	Website  string
	Email    string
}

func GetUserById(id int) User {
	client := &http.Client{}
	userUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/users/", id)
	req, err := http.NewRequest(http.MethodGet, userUrl, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	var resObject User
	json.Unmarshal(bodyBytes, &resObject)
	return resObject
}
