package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"pr/utiles"
	"pr/models"
)

const numWorkers = 3
const redisListName = "myList"

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	ch := make(chan models.TaskSignal, numWorkers)
	go models.DoWork(ch)
	go listenRedis(ch)
	r.Run()
}

func regTask(ch chan models.Task) {
	//taskList := models.TaskList{}
	//taskList.AddTask(models.Task{Name: "push article"})
	//taskList.AddTask(models.Task{Name: "edit article"})
}

func listenRedis(ch chan models.TaskSignal) {
	client := utiles.RedisClient
	defer client.Close()
	channelStrings := []string{"NJ", "SH"}
	psc := client.Subscribe(channelStrings...)
	for {
		v, _ := psc.ReceiveMessage()
		ch <- models.TaskSignal{Start: 1}
		client.RPush(redisListName, v.Payload)
		fmt.Println(v.Payload)
	}
}
