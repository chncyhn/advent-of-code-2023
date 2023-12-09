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

type Digit struct {
	colStart int
	colEnd   int
	row      int
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func areAdj(dig Digit, symbol Coord) bool {
	if abs(symbol.x-dig.row) > 1 {
		return false
	}
	if symbol.y < dig.colStart {
		return abs(symbol.y-dig.colStart) <= 1
	} else if symbol.y > dig.colEnd {
		return abs(symbol.y-dig.colEnd) <= 1
	}
	return true
}

func digValue(digit Digit, board []string) int {
	return extractDig(board, digit.row, digit.colStart, digit.colEnd)
}

func part1(board []string) (digs []Digit, symbols []Coord) {
	isAdj := make([][]bool, len(board))
	for i := range isAdj {
		isAdj[i] = make([]bool, len(board[0]))
	}
	n, m := len(board), len(board[0])
	for i := range board {
		for j := range board[i] {
			if isSymbol(board[i][j]) {
				symbols = append(symbols, Coord{i, j})
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
				digs = append(digs, Digit{digStart, j - 1, i})
				digStart = -1
			}
		}
		if digStart != -1 && anyAdj(isAdj, i, digStart, len(row)-1) {
			ans += extractDig(board, i, digStart, len(row)-1)
			digs = append(digs, Digit{digStart, len(row) - 1, i})
		}
	}
	fmt.Println(ans)
	return
}

func part2(symbols []Coord, digs []Digit, board []string) {
	adjDigsPerSymbol := make(map[Coord][]Digit)
	for _, symbol := range symbols {
		for _, dig := range digs {
			if areAdj(dig, symbol) {
				adjDigsPerSymbol[symbol] = append(adjDigsPerSymbol[symbol], dig)
			}
		}
	}
	ans := 0
	for _, adjDigs := range adjDigsPerSymbol {
		if len(adjDigs) == 2 {
			ans += digValue(adjDigs[0], board) * digValue(adjDigs[1], board)
		}
	}
	fmt.Println(ans)
}

func main() {
	board := readInput()
	digs, symbols := part1(board)
	part2(symbols, digs, board)
}
