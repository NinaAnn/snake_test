package component

import (
	"fmt"
	"math/rand"
	"time"
)

func NewPannel(w int, h int) *Pannel {
	r := &Pannel{
		totalCount: 0,
		width:      w,
		height:     h,
		snakeMap:   make(map[int64]*Snake),
	}
	fmt.Println("add pannel, pannel width:" + string(r.width))

	var pan [][]int
	for i := 0; i < h; i++ {
		inline := make([]int, w)
		pan = append(pan, inline)
	}
	r.pan = pan
	go r.Run()

	return r
}

type Pannel struct {
	totalCount int
	width      int
	height     int
	snakeMap   map[int64]*Snake
	pan        [][]int
}

func (P *Pannel) Enter(name string) *Snake {
	var x int
	var y int
	for {
		x = rand.Intn(P.width)
		y = rand.Intn(P.height)
		if P.pan[x][y] != 1 {
			break
		}
	}
	n := time.Now().Unix()
	snake := NewSnake(P.width, P.height, name, x, y, n)
	P.snakeMap[n] = snake
	P.totalCount += 1
	return snake
}

func (P *Pannel) Exit(id int64) {
	snake := P.snakeMap[id]

}

func (P *Pannel) Run() {
	for {
		P.Update()
		for _, v := range P.snakeMap {
			var pan [][]int
			for i := 0; i < P.height; i++ {
				pan = append(pan, P.pan[i])
			}
			pan[v.X][v.Y] = 3
			v.PosChan <- pan
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func (P *Pannel) Update() {
	fmt.Println("update:")
	var pan [][]int
	for i := 0; i < P.height; i++ {
		inline := make([]int, P.width)
		pan = append(pan, inline)
	}

	for _, snake := range P.snakeMap {
		_, err := snake.Move()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, v := range snake.pos {
			pan[v[0]][v[1]] = 1
		}
	}
	for _, snake := range P.snakeMap {
		x := snake.X
		y := snake.Y
		if pan[x][y] == 1 {
			return
		}
		pan[x][y] = 2

	}
	fmt.Println(pan)
	P.pan = P.pan[0:0][0:0]
	for i := 0; i < P.height; i++ {
		P.pan = append(P.pan, pan[i])
	}

}
