package main

import (
	"fmt"
	"gameoflife/game"
	"math/rand"
	"time"
)

var gridSize = 40

func main() {
	arr := make([][]int, gridSize)
	for i := range arr {
		arr[i] = make([]int, gridSize)
	}
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			src := rand.NewSource(time.Now().UnixNano())
			r := rand.New(src)
			arr[i][j] = r.Intn(2)
		}
	}
	fmt.Print("\033[H\033[2J")

	generation := 1

	for {

		population := 0
		generation++

		arr2 := make([][]int, gridSize)
		for i := range arr2 {
			arr2[i] = make([]int, gridSize)
		}
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				population += arr[j][i]
				neighbours := game.CountNeighbours(arr, game.Position{Row: j, Column: i})
				arr2[j][i] = game.GetNextValueByNeighboursCount(arr[j][i], neighbours)
			}
		}
		game.Draw(arr)
		fmt.Printf("Generation %d \n", generation)
		fmt.Printf("Population %d \n", population)

		arr = arr2

		time.Sleep(200 * time.Millisecond)

		fmt.Print("\033[H\033[2J")

	}
}
