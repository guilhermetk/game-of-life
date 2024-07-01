package main

import (
	"fmt"
	"math/rand"
	"time"
)

var gridSize = 20

func main() {
	arr := make([][]int, gridSize)
	for i := range arr {
		arr[i] = make([]int, gridSize)
	}
	fmt.Print("\033[H\033[2J")
	for {
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				src := rand.NewSource(time.Now().UnixNano())
				r := rand.New(src)
				arr[i][j] = r.Intn(2)
			}
		}
		draw(arr)

		time.Sleep(1000 * time.Millisecond)

		fmt.Print("\033[H\033[2J")

	}
}

func draw(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d|", arr[i][j])
		}
		fmt.Println()
	}
}
