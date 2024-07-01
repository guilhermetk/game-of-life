package game

import "fmt"

type Position struct {
	Row    int
	Column int
}

func Draw(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[j][i] == 1 {
				fmt.Printf("%s|", "X")
			} else {
				fmt.Printf("%s|", " ")
			}
		}
		fmt.Println()
	}
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

func GetNextValueByNeighboursCount(currentCellState int, cellNeighboursCount int) int {
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
		return 0
	}
}
