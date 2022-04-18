package todo

import (
	"encoding/json"
	"fmt"
	"github.com/ingkiller/hackernews/internal/client"
	"io/ioutil"
	"sort"
)

type Task struct {
	Id        int
	UserId    int
	Title     string
	Completed bool
}

var TaskByUserId = make(map[int]map[int]*Task)
var nextTask = 1

func store(task Task, userId int) {
	if TaskByUserId[userId] == nil {
		TaskByUserId[userId] = make(map[int]*Task)
	}
	TaskByUserId[userId][task.Id] = &task
	nextTask++
}

func GetListByUserId(userId int) []Task {
	var result []Task

	if len(TaskByUserId[userId]) > 0 {
		for _, t := range TaskByUserId[userId] {
			result = append(result, *t)
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i].Id < result[j].Id
		})
		return result
	}

	userUrl := fmt.Sprint("https://jsonplaceholder.typicode.com/todos?userId=", userId)

	resp, err := client.MakeReq(userUrl)
	if err != nil {
		fmt.Print("NewRequest: %v", err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &result)

	for _, task := range result {
		store(task, userId)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Id < result[j].Id
	})
	return result
}

func ToggleTask(taskId int, userId int) Task {
	task := TaskByUserId[userId][taskId]
	task.Completed = !task.Completed
	TaskByUserId[userId][taskId] = task
	return *task
}

func CreateTask(title string, userId int) Task {
	newTask := Task{
		Id:        nextTask,
		UserId:    userId,
		Title:     title,
		Completed: false,
	}
	nextTask++
	store(newTask, userId)
	return newTask
}

func DeleteTask(taskId int, userId int) {
	delete(TaskByUserId[userId], taskId)
}
