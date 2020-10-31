package server

import (
	"snake_test/component"

	"github.com/gin-gonic/gin"
)

var room = component.NewRoom()

func NewServer() *gin.Engine {
	s := gin.Default()
	{
		// websocket
		s.GET("/ws/socket", Websocket.Handle())
	}
	return s
}
