package server

import (
	"github.com/gin-gonic/gin"
	"snake_ini/component"
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
