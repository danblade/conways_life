package main

// asked bing ai chat: https://sl.bing.net/eGDsPlEMcAC

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 40
	height = 20
)

type World [][]bool

func NewWorld() World {
	w := make(World, height)
	for i := range w {
		w[i] = make([]bool, width)
	}
	return w
}

func (w World) Show() {
	for _, row := range w {
		for _, cell := range row {
			if cell {
				fmt.Print("■ ")
			} else {
				fmt.Print("□ ")
			}
		}
		fmt.Println()
	}
}

func (w World) Next() World {
	nw := NewWorld()
	for i, row := range w {
		for j := range row {
			nw[i][j] = w[i][j]
			n := w.Neighbors(i, j)
			if n < 2 || n > 3 {
				nw[i][j] = false
			} else if n == 3 {
				nw[i][j] = true
			}
		}
	}
	return nw
}

func (w World) Neighbors(i, j int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			ni, nj := i+x, j+y
			if ni < 0 || nj < 0 || ni >= height || nj >= width {
				continue
			}
			if w[ni][nj] {
				count++
			}
		}
	}
	return count
}

func main() {
	rand.Seed(time.Now().UnixNano())
	world := NewWorld()
	for i := range world {
		for j := range world[i] {
			world[i][j] = rand.Intn(2) == 1
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Generation %d:\n", i+1)
		world.Show()
		world = world.Next()
		time.Sleep(500 * time.Millisecond)
	}
}
