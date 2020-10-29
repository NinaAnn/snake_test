package component

import (
	"fmt"
	"math/rand"
)

func NewRoom() *Room {
	fmt.Println("create room")
	r := &Room{
		totalCount: 1,
		pannelMap:  make(map[int]*Pannel),
		userMap:    make(map[int64]int),
		joinChan:   make(chan int64),
		leaveChan:  make(chan string),
	}
	w := 20
	h := 20
	i := 0
	for i < r.totalCount {
		r.pannelMap[i] = NewPannel(w, h)
		i += 1
	}
	fmt.Println("room create finished")
	return r
}

type Room struct {
	totalCount int
	pannelMap  map[int]*Pannel
	userMap    map[int64]int
	joinChan   chan int64
	leaveChan  chan string
}

func (R *Room) Exit(id int64) {
	roomI := R.userMap[id]
	pannel := R.pannelMap[roomI]
	pannel.Exit(id)
	delete(R.userMap, id)
	return
}

func (R *Room) Enter(name string) *Snake {
	roomI := rand.Intn(1)
	pannel := R.pannelMap[roomI]
	u := pannel.Enter(name)
	R.userMap[u.Id] = roomI
	return u
}
