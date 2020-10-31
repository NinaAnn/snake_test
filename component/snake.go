package component

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var DirMap = map[int][2]int{
	0: [2]int{0, 1},
	1: [2]int{0, -1},
	2: [2]int{1, 0},
	3: [2]int{-1, 0},
}

var ConditionMap = map[int]map[int]int{
	0: map[int]int{
		0: 0,
		1: 0,
		2: 2,
		3: 3,
	},
	1: map[int]int{
		0: 1,
		1: 1,
		2: 3,
		3: 2,
	},
	2: map[int]int{
		0: 2,
		1: 2,
		2: 1,
		3: 0,
	},
	3: map[int]int{
		0: 3,
		1: 3,
		2: 0,
		3: 1,
	},
}

func NewSnake(w int, h int, name string, x int, y int, n int64) *Snake {
	rand.Seed(time.Now().Unix())
	s := &Snake{
		w:        w,
		h:        h,
		X:        x,
		Y:        y,
		Name:     name,
		Length:   2,
		MoveChan: make(chan int),
		PosChan:  make(chan [][]int),
		pos:      [][2]int{},
		Id:       n,
		DoneChan: make(chan int),
		ExitChan: make(chan string),
	}
	d := rand.Intn(3)
	s.dir = d
	s.Add()
	go s.Serve()
	return s
}

type Snake struct {
	w        int
	h        int
	X        int
	Y        int
	dir      int
	Length   int
	Name     string
	MoveChan chan int
	pos      [][2]int
	Id       int64
	PosChan  chan [][]int
	DoneChan chan int
	ExitChan chan string
}

func (s *Snake) Serve() {
	for {
		select {
		// 执行操作
		case id := <-s.MoveChan:
			s.dir = ConditionMap[s.dir][id]
			fmt.Println(s.dir)
		case <-s.DoneChan:
			fmt.Println("exiting...")
			s.DoneChan <- 1
			break
		default:
		}
	}
}

func (s *Snake) Move() (n int, err error) {
	insert := [2]int{s.X, s.Y}
	dir := DirMap[s.dir]
	s.X += dir[0]
	s.Y += dir[1]
	if s.X >= s.w || s.X < 0 || s.Y < 0 || s.Y >= s.h {
		return 1, errors.New("snake out of range")
	}
	var newPos [][2]int
	newPos = append(newPos, insert)
	if len(s.pos) > 0 {
		s.pos = append(newPos, s.pos[0:(len(s.pos)-1)]...)
	} else {
		s.pos = newPos
	}
	return 0, nil
}

func (s *Snake) Add() {
	var tail [2]int
	if len(s.pos) > 0 {
		tail = s.pos[len(s.pos)-1]
	} else {
		tail = [2]int{s.X, s.Y}
	}
	before := [2]int{-1, -1}
	if len(s.pos) > 1 {
		before = s.pos[len(s.pos)-2]
	}
	candidates := [4][2]int{{tail[0], tail[1] + 1}, {tail[0], tail[1] - 1}, {tail[0] + 1, tail[1]}, {tail[0] - 1, tail[1]}}
	var newTail [2]int
	for _, v := range candidates {
		if v[0] == s.X && v[1] == s.Y {
			continue
		}
		if v[0] == before[0] && v[1] == before[1] {
			continue
		}
		if v[0] >= s.w || v[0] < 0 || v[1] < 0 || v[1] >= s.h {
			continue
		}
		newTail = v
		break
	}
	s.pos = append(s.pos, newTail)
}
