package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Task struct {
	Id        int
	UserId    int
	Title     string
	Completed bool
}

var TasksById = make(map[int]*Task)
var nextTask = 1

func store(task Task) {
	TasksById[task.Id] = &task
	nextTask++
}

func GetListByUserId(userId int) []Task {

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
	var result []Task
	json.Unmarshal(bodyBytes, &result)

	for _, task := range result {
		store(task)
	}
	return result
}

func ToggleTask(taskId int) Task {
	task := TasksById[taskId]
	task.Completed = !task.Completed
	TasksById[taskId] = task
	return *task
}

func CreateTask(title string) Task {
	newTask := Task{
		Id:        nextTask,
		UserId:    1,
		Title:     title,
		Completed: false,
	}
	nextTask++
	store(newTask)
	return newTask
}

func DeleteTask(userId int) {
	delete(TasksById, userId)
}
