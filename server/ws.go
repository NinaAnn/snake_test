package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var Websocket = &ws{
	upgrader: &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	},
}

type ws struct {
	upgrader *websocket.Upgrader
}

func newWebSocket() *ws {
	ws := new(ws)
	ws.upgrader = &websocket.Upgrader{
		ReadBufferSize:  4096, //指定读缓存区大小
		WriteBufferSize: 1024, // 指定写缓存区大小
		// 检测请求来源
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("method is not GET")
				return false
			}
			if r.URL.Path != "/ws" {
				fmt.Println("path error")
				return false
			}
			return true
		},
	}
	return ws
}

func (s *ws) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
		defer func() {
			conn.Close()
		}()
		if err != nil {
			msg := map[string]string{"code": "1", "msg": "error"}
			conn.WriteJSON(msg)
			return
		}
		u := room.Enter(name)
		retMsg := make(map[string]interface{})

		for {
			select {
			case data := <-u.PosChan:
				retMsg["name"] = name
				retMsg["code"] = "0"
				retMsg["data"] = data
				err := conn.WriteJSON(retMsg)
				if err != nil {
					fmt.Println("send msg faild ", err)
					return
				}
			}
		}
		return
	}
}
