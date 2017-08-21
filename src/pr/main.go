package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"pr/utiles"
)

type Task struct {
	name string
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//numWorkers := 3
	//ch := make(chan Task, 3)
	//// Run fixed number of workers.
	//for i := 0; i < numWorkers; i++ {
	//	go worker(ch)
	//}
	//// Send tasks to workers.
	//Tasks := getTasks()
	//for _, task := range Tasks {
	//	fmt.Println("range task list", task)
	//	ch <- task
	//}
	go listenRedis([]string{"hhh", "mmm"})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func regTask() {

}

func listenRedis(task []string) {
	client := utiles.RedisClient
	defer client.Close()
	fmt.Println(task)
	psc := client.Subscribe(task...)
	for {
		v, _ := psc.ReceiveMessage()
		fmt.Println(v.Payload)
	}
}

func worker(ch chan Task) {
	fmt.Println("start one task worker")
	for {
		task := <-ch
		process(task)
	}
}

func getTasks() []Task {
	all := []Task{}
	task1 := Task{name: "a"}
	task2 := Task{name: "b"}
	return append(all, task1, task2)
}

func process(ch Task) {
	fmt.Println("complete this task", ch.name)
}
