package main

import (
	"bufio"
	"os"
)

type Grid [][]rune

func allValid(int, int) bool { return true }

func readGrids() (grids []Grid) {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	var curGrid Grid
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			grids = append(grids, curGrid)
			curGrid = Grid{}
		} else {
			row := []rune{}
			for _, c := range txt {
				row = append(row, c)
			}
			curGrid = append(curGrid, row)
		}
	}
	grids = append(grids, curGrid)
	return
}

func isMirroringFromRow(grid Grid, row int) bool {
	for o := 0; row-o >= 0 && row+o+1 < len(grid); o++ {
		i := row - o
		j := row + o + 1
		for k := 0; k < len(grid[i]); k++ {
			if grid[i][k] != grid[j][k] {
				return false
			}
		}
	}
	return true
}

func isMirroringFromCol(grid Grid, col int) bool {
	for o := 0; col-o >= 0 && col+o+1 < len(grid[0]); o++ {
		i := col - o
		j := col + o + 1
		for k := 0; k < len(grid); k++ {
			if grid[k][i] != grid[k][j] {
				return false
			}
		}
	}
	return true
}

func findMirror(grid Grid, valid func(int, int) bool) (int, int) {
	for i := 0; i < len(grid)-1; i++ {
		if isMirroringFromRow(grid, i) && valid(i, -1) {
			return i, -1
		}
	}
	for i := 0; i < len(grid[0])-1; i++ {
		if isMirroringFromCol(grid, i) && valid(-1, i) {
			return -1, i
		}
	}
	return -1, -1
}

func mirrorToScore(row, col int) int {
	if row != -1 {
		return 100 * (row + 1)
	} else if col != -1 {
		return (col + 1)
	} else {
		panic("no mirror")
	}
}

func flip(r rune) rune {
	if r == '#' {
		return '.'
	} else if r == '.' {
		return '#'
	}
	panic("invalid rune")
}

func search(grid Grid) (int, int) {
	baseRow, baseCol := findMirror(grid, allValid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = flip(grid[i][j])

			newRow, newCol := findMirror(grid, func(i, j int) bool { return i != baseRow || j != baseCol })
			if (newRow != -1 || newCol != -1) && ((newRow != baseRow) || (newCol != baseCol)) {
				return newRow, newCol
			}

			grid[i][j] = flip(grid[i][j])
		}
	}
	panic("no solution")
}

func part1(grids []Grid) {
	ans := 0
	for _, grid := range grids {
		row, col := findMirror(grid, allValid)
		ans += mirrorToScore(row, col)
	}
	println(ans)
}

func part2(grids []Grid) {
	ans := 0
	for _, grid := range grids {
		row, col := search(grid)
		ans += mirrorToScore(row, col)
	}
	println(ans)
}

func main() {
	grids := readGrids()
	part1(grids)
	part2(grids)
}
