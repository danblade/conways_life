package main

import (
	"fmt"
	"time"

	"github.com/danblade/conways_life/cgol"
)

func main() {
	game := cgol.NewGame(15, 15)
	game.Set(3, 3, true)
	game.Set(1, 3, true)
	game.Set(3, 2, true)
	game.Set(5, 1, true)
	game.Set(2, 3, true)
	game.Set(4, 3, true)
	game.Set(13, 2, true)
	game.Set(12, 1, true)
	for {
		render(game)
		game.Step()
		time.Sleep(time.Second)
		fmt.Println("------------------------------------------")
	}
}

func render(game *cgol.Game) {
	for y := 0; y < game.Height(); y++ {
		for x := 0; x < game.Width(); x++ {
			cell := game.Get(x, y)
			if cell {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
