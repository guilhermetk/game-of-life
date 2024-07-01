package main

import "gameoflife/game"

func main() {
	game := game.New(40)
	game.Run()
}
