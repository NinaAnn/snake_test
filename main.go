package main

import (
	"log"
	"math/rand"
	"snake_test/server"
	"time"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
	rand.Seed(time.Now().UnixNano())
	s := server.NewServer()
	log.Fatal(s.Run())
}
