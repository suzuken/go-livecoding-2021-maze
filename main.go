package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Maze struct {
	width, height int
	cells         [][]int
}

func (m *Maze) set(x, y, v int) {
	m.cells[y][x] = v
}

func (m *Maze) inMap(x, y int) bool {
	return x >= 0 && x <= m.width-1 && y >= 0 && y <= m.height-1
}

func (m *Maze) get(x, y int) int {
	if !m.inMap(x, y) {
		return -1
	}
	return m.cells[y][x]
}

type cell struct {
	x, y int
}

func (m *Maze) init(width, height int) {
	m.width = width
	m.height = height

	m.cells = make([][]int, height)
	for i := 0; i < height; i++ {
		m.cells[i] = make([]int, width)
		for j := 0; j < width; j++ {
			m.cells[i][j] = 1 // 壁
		}
	}
}

func (m *Maze) printMap() {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			v := m.cells[i][j]
			if v == 0 {
				fmt.Print("　")
				continue
			}
			if v == 1 {
				fmt.Print("壁")
			}
		}
		fmt.Println("")
	}
}

func (m *Maze) Generate() {
	start := cell{rand.Intn(m.width), rand.Intn(m.height)}
	fmt.Printf("start = %+v\n", start)

	stack := make([]cell, 0)
	stack = append(stack, start)
	m.set(start.x, start.y, 0)

	var cur cell
	cur = start

	candidates := func(c cell) []cell {
		ret := make([]cell, 0)
		// 上下左右
		if m.get(c.x, c.y-1) == 1 && m.get(c.x, c.y-2) != 0 && m.get(c.x-1, c.y-1) != 0 && m.get(c.x+1, c.y-1) != 0 {
			ret = append(ret, cell{c.x, c.y - 1}) // 上
		}
		if m.get(c.x, c.y+1) == 1 && m.get(c.x, c.y+2) != 0 && m.get(c.x-1, c.y+1) != 0 && m.get(c.x+1, c.y+1) != 0 {
			ret = append(ret, cell{c.x, c.y + 1}) // 下
		}
		if m.get(c.x-1, c.y) == 1 && m.get(c.x-2, c.y) != 0 && m.get(c.x-1, c.y+1) != 0 && m.get(c.x-1, c.y-1) != 0 {
			ret = append(ret, cell{c.x - 1, c.y}) // 左
		}
		if m.get(c.x+1, c.y) == 1 && m.get(c.x+2, c.y) != 0 && m.get(c.x+1, c.y+1) != 0 && m.get(c.x+1, c.y-1) != 0 {
			ret = append(ret, cell{c.x + 1, c.y}) // 右
		}
		return ret
	}

	for len(stack) >= 1 {
		for {
			// 方向をそれぞれみるー
			cc := candidates(cur)
			if len(cc) == 0 {
				break
			}
			// 全方向だめならぬける
			// どっちにいくかきめて
			// そのマスを道にする
			// つぎにいく・・
			next := cc[rand.Intn(len(cc))]
			stack = append(stack, next)
			m.set(next.x, next.y, 0)
			cur = next
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	m := Maze{}
	m.init(7, 7)
	m.printMap()
	fmt.Printf("m = %+v\n", m)
	m.Generate()
	m.printMap()
}
