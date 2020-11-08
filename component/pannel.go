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
		foodCount:  (int)(w / 2),
		foodMap:    make(map[[2]int]int),
	}
	fmt.Println("add pannel, pannel width:" + string(r.width))
	r.initFood()
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
	foodCount  int
	foodMap    map[[2]int]int
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
	delete(P.snakeMap, id)
	snake.DoneChan <- 1
	<-snake.DoneChan
	snake = nil
	P.totalCount--
}

func (P *Pannel) initFood() {
	i := 0
	for {
		x := rand.Intn(P.width)
		y := rand.Intn(P.height)
		if _, ok := P.foodMap[[2]int{x, y}]; ok {
			continue
		} else {
			P.foodMap[[2]int{x, y}] = 1
			i++
		}
		if i == P.foodCount {
			break
		}
	}
}

func (P *Pannel) addFood(count int) {
	i := 0
	for {
		if i == count {
			break
		}
		x := rand.Intn(P.width)
		y := rand.Intn(P.height)
		if _, ok := P.foodMap[[2]int{x, y}]; ok {
			continue
		} else {
			if P.pan[x][y] == 2 {
				continue
			} else {
				P.foodMap[[2]int{x, y}] = 1
				i++
			}
		}
	}
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
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func (P *Pannel) Update() {
	fmt.Println("update:")
	addTag := 0
	var pan [][]int
	for i := 0; i < P.height; i++ {
		inline := make([]int, P.width)
		pan = append(pan, inline)
	}

	for _, snake := range P.snakeMap {
		_, err := snake.Move()
		if err != nil {
			fmt.Println(err)
			snake.ExitChan <- err.Error()
			P.Exit(snake.Id)
		}
		for _, v := range snake.pos {
			pan[v[0]][v[1]] = 1
		}
	}
	for _, snake := range P.snakeMap {
		x := snake.X
		y := snake.Y
		if pan[x][y] == 1 || pan[x][y] == 2 {
			snake.ExitChan <- "sorry you are out of the game"
			P.Exit(snake.Id)
			pan[x][y] = 2
			continue
		}
		if _, ok := P.foodMap[[2]int{x, y}]; ok {
			delete(P.foodMap, [2]int{x, y})
			snake.Add()
			addTag++
		}
	}
	// fmt.Println(pan)
	P.pan = P.pan[0:0][0:0]
	for i := 0; i < P.height; i++ {
		P.pan = append(P.pan, pan[i])
	}
	P.addFood(addTag)
	for pos := range P.foodMap {
		P.pan[pos[0]][pos[1]] = 10
	}
}
