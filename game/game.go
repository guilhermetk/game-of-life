package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	grid [][]int
}

type Position struct {
	Row    int
	Column int
}

func New(gridSize int) *Game {
	g := Game{}
	g.grid = make([][]int, gridSize)
	for i := range g.grid {
		g.grid[i] = make([]int, gridSize)
	}
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			src := rand.NewSource(time.Now().UnixNano())
			r := rand.New(src)
			g.grid[i][j] = r.Intn(2)
		}
	}
	return &g
}

func (g *Game) Run() {
	gridSize := len(g.grid[0])
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
				population += g.grid[j][i]
				neighbours := CountNeighbours(g.grid, Position{Row: j, Column: i})
				arr2[j][i] = getNextValueByNeighboursCount(g.grid[j][i], neighbours)
			}
		}
		draw(g.grid, generation, population)
		g.grid = arr2
		time.Sleep(400 * time.Millisecond)
	}
}

func draw(arr [][]int, generation, population int) {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[j][i] == 1 {
				fmt.Printf("%s ", "X")
			} else {
				fmt.Printf("%s ", " ")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Generation: %d \n", generation)
	fmt.Printf("Population: %d \n", population)
}

var neighbours = []Position{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func CountNeighbours(grid [][]int, p Position) int {
	count := 0
	for _, n := range neighbours {
		neighbourRow := p.Row + n.Row
		neighbourColumn := p.Column + n.Column

		if neighbourRow >= 0 && neighbourRow < len(grid) && neighbourColumn >= 0 && neighbourColumn < len(grid[0]) && grid[neighbourRow][neighbourColumn] > 0 {
			count++
		}
	}
	return count
}

func getNextValueByNeighboursCount(currentCellState int, cellNeighboursCount int) int {
	switch {
	case currentCellState == 1 && cellNeighboursCount < 2:
		return 0
	case currentCellState == 1 && (cellNeighboursCount == 2 || cellNeighboursCount == 3):
		return 1
	case currentCellState == 1 && cellNeighboursCount > 3:
		return 0
	case currentCellState == 0 && cellNeighboursCount == 3:
		return 1
	default:
		return currentCellState
	}
}
