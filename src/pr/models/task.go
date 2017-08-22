package models

import "fmt"

type TaskList struct {
	List []Task
}

type Task struct {
	Name string
}

func (t *TaskList) AddTask(tl Task) TaskList {
	t.List = append(t.List, tl)
	return *t
}

func DoWork(ch chan Task) {
	fmt.Println("start one task worker")
	for {
		task := <-ch
		task.process()
	}
}

func (t *Task) process() {
	fmt.Println("do this task", t.Name)
}
