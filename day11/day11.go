package main

import (
	"bufio"
	"os"
)

func readInput(fl string) (board [][]int) {
	file, _ := os.Open(fl)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		for _, ch := range scanner.Text() {
			if ch == '#' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		board = append(board, row)
	}
	return
}

type Coord struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattan(a, b Coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

func dist(a, b Coord, zeroRows, zeroCols map[int]bool, multiplier int) int {
	extras := 0
	for i := min(a.x, b.x) + 1; i < max(a.x, b.x); i++ {
		if zeroRows[i] {
			extras++
		}
	}
	for j := min(a.y, b.y) + 1; j < max(a.y, b.y); j++ {
		if zeroCols[j] {
			extras++
		}
	}
	return manhattan(a, b) + extras*(multiplier-1)
}

func allZeroes(series []int) bool {
	for _, v := range series {
		if v != 0 {
			return false
		}
	}
	return true
}

func column(board [][]int, col int) (series []int) {
	for _, row := range board {
		series = append(series, row[col])
	}
	return
}

func main() {
	board := readInput("full.txt")
	rows, cols := len(board), len(board[0])
	zeroRows, zeroCols := make(map[int]bool), make(map[int]bool)
	for i := 0; i < rows; i++ {
		if allZeroes(board[i]) {
			zeroRows[i] = true
		}
	}
	for j := 0; j < cols; j++ {
		if allZeroes(column(board, j)) {
			zeroCols[j] = true
		}
	}

	var astreoids []Coord
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 1 {
				astreoids = append(astreoids, Coord{i, j})
			}
		}
	}

	part1 := 0
	part2 := 0
	for i := 0; i < len(astreoids); i++ {
		for j := i + 1; j < len(astreoids); j++ {
			part1 += dist(astreoids[i], astreoids[j], zeroRows, zeroCols, 2)
			part2 += dist(astreoids[i], astreoids[j], zeroRows, zeroCols, 1_000_000)
		}
	}
	println(part1)
	println(part2)
}
