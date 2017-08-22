package models

import (
	"fmt"
	"pr/utiles"
	"encoding/json"
)

type TaskSignal struct {
	Start int
}

type TaskList struct {
	List []Task
}

type Task struct {
	Name string `json:"name"`
}

func (t *TaskList) AddTask(tl Task) TaskList {
	t.List = append(t.List, tl)
	return *t
}

func DoWork(ch chan TaskSignal) {
	fmt.Println("start one task worker")
	client := utiles.RedisClient
	defer client.Close()
	for {
		<-ch
		pop, err := client.LPop("myList").Result()
		fmt.Println(err)
		var task Task
		json.Unmarshal([]byte(pop), &task)
		fmt.Println("pop ", pop)
		fmt.Println("&task ", task.Name)
		task.process()
	}
}

func (t *Task) process() {
	fmt.Println("do this task", t.Name)
}
