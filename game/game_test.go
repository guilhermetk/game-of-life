package game

import "testing"

func TestNeighbours(t *testing.T) {
	t.Run("test neighbours count", func(t *testing.T) {
		grid := [][]int{
			{0, 0, 0, 1, 0},
			{0, 1, 0, 0, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0}}

		testPosition := Position{Row: 1, Column: 2}
		neighboursCount := CountNeighbours(grid, testPosition)
		expectedCound := 4

		if neighboursCount != expectedCound {
			t.Errorf("expected %d neighbours, got %d", expectedCound, neighboursCount)
		}
	})

	t.Run("test neighbours count on edges", func(t *testing.T) {
		grid := [][]int{
			{0, 1, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0}}

		testPosition := Position{Row: 0, Column: 0}
		neighboursCount := CountNeighbours(grid, testPosition)
		expectedCound := 3

		if neighboursCount != expectedCound {
			t.Errorf("expected %d neighbours, got %d", expectedCound, neighboursCount)
		}
	})

	t.Run("test neighbours count on edges", func(t *testing.T) {
		grid := [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 1, 0}}

		testPosition := Position{Row: 4, Column: 4}
		neighboursCount := CountNeighbours(grid, testPosition)
		expectedCound := 3

		if neighboursCount != expectedCound {
			t.Errorf("expected %d neighbours, got %d", expectedCound, neighboursCount)
		}
	})
}
