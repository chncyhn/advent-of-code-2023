package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput() (board []string) {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		board = append(board, scanner.Text())
	}
	return
}

type Coord struct {
	x, y int
}

func neighs(c Coord, n, m int) []Coord {
	var neighs []Coord
	for i := c.x - 1; i <= c.x+1; i++ {
		for j := c.y - 1; j <= c.y+1; j++ {
			if i >= 0 && i < n && j >= 0 && j < m && !(i == c.x && j == c.y) {
				neighs = append(neighs, Coord{i, j})
			}
		}
	}
	return neighs
}

func isSymbol(r byte) bool {
	return r != '.' && (r < '0' || r > '9')
}

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func anyAdj(isAdj [][]bool, row, digStart, digEnd int) bool {
	for i := digStart; i <= digEnd; i++ {
		if isAdj[row][i] {
			return true
		}
	}
	return false
}

func extractDig(board []string, row, digStart, digEnd int) int {
	dig := 0
	for i := digStart; i <= digEnd; i++ {
		dig = dig*10 + int(board[row][i]-'0')
	}
	return dig
}

func main() {
	board := readInput()

	isAdj := make([][]bool, len(board))
	for i := range isAdj {
		isAdj[i] = make([]bool, len(board[0]))
	}

	n, m := len(board), len(board[0])
	for i := range board {
		for j := range board[i] {
			if isSymbol(board[i][j]) {
				for _, neigh := range neighs(Coord{i, j}, n, m) {
					isAdj[neigh.x][neigh.y] = true
				}
			}
		}
	}

	ans := 0
	for i, row := range board {
		digStart := -1
		for j, col := range row {
			if isDigit(byte(col)) && digStart == -1 {
				digStart = j
			} else if !isDigit(byte(col)) && digStart != -1 {
				if anyAdj(isAdj, i, digStart, j-1) {
					ans += extractDig(board, i, digStart, j-1)
				}
				digStart = -1
			}
		}
		if digStart != -1 && anyAdj(isAdj, i, digStart, len(row)-1) {
			ans += extractDig(board, i, digStart, len(row)-1)
		}
	}
	fmt.Println(ans)
}
