package main

import (
	"bufio"
	"os"
)

func readInput() (ret [][]byte) {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, []byte(scanner.Text()))
	}
	return
}

type Coord struct {
	x, y int
}

func possibles(grid [][]byte, from Coord) []Coord {
	cur := grid[from.x][from.y]
	if cur == '-' {
		return []Coord{{from.x, from.y + 1}, {from.x, from.y - 1}}
	}
	if cur == '|' {
		return []Coord{{from.x + 1, from.y}, {from.x - 1, from.y}}
	}
	if cur == 'L' {
		return []Coord{{from.x, from.y + 1}, {from.x - 1, from.y}}
	}
	if cur == 'J' {
		return []Coord{{from.x, from.y - 1}, {from.x - 1, from.y}}
	}
	if cur == '7' {
		return []Coord{{from.x, from.y - 1}, {from.x + 1, from.y}}
	}
	if cur == 'F' {
		return []Coord{{from.x, from.y + 1}, {from.x + 1, from.y}}
	}
	return []Coord{}
}

func nextPos(grid [][]byte, from Coord) (ret []Coord) {
	for _, v := range possibles(grid, from) {
		if grid[v.x][v.y] == 'S' {
			ret = append(ret, v)
			continue
		}
		for _, v2 := range possibles(grid, v) {
			if v2 == from {
				ret = append(ret, v)
				break
			}
		}
	}
	return
}

func findStart(grid [][]byte) Coord {
	for x, line := range grid {
		for y, char := range line {
			if char == 'S' {
				return Coord{x, y}
			}
		}
	}
	panic("no start found")
}

func allNeighs(from Coord, grid [][]byte) (ret []Coord) {
	for _, cand := range []Coord{{from.x + 1, from.y}, {from.x - 1, from.y}, {from.x, from.y + 1}, {from.x, from.y - 1}} {
		if cand.x >= 0 && cand.x < len(grid) && cand.y >= 0 && cand.y < len(grid[0]) {
			ret = append(ret, cand)
		}
	}
	return
}

func findSuitableConn(grid [][]byte, start Coord) Coord {
	for _, conn := range allNeighs(start, grid) {
		for _, v := range possibles(grid, conn) {
			if v == start {
				return conn
			}
		}
	}
	panic("no suitable conn found")
}

func findLoop(grid [][]byte) (path []Coord) {
	start := findStart(grid)
	path = append(path, start)
	cur := findSuitableConn(grid, start)
	visited := make(map[Coord]bool)
	visited[start] = true
	visited[cur] = true
	scnt := 0
	for scnt < 2 {
		path = append(path, cur)
		np := nextPos(grid, cur)
		for _, v := range np {
			if grid[v.x][v.y] == 'S' {
				scnt++
			}
			if _, ok := visited[v]; !ok {
				cur = v
				visited[cur] = true
				break
			}
		}
	}
	return
}

func isInLoop(point Coord, loop []Coord) bool {
	intersections := 0
	n := len(loop)
	for i := 0; i < n; i++ {
		v0 := loop[i]
		v1 := loop[(i+1)%n]
		if (point.y > v0.y) != (point.y > v1.y) &&
			point.x < (v1.x-v0.x)*(point.y-v0.y)/(v1.y-v0.y)+v0.x {
			intersections++
		}
	}
	return intersections%2 == 1
}

func dfs(from Coord, grid [][]byte, visited map[Coord]bool, loop []Coord) int {
	if _, ok := visited[from]; ok {
		return 0
	}
	visited[from] = true
	queue := []Coord{from}
	allInLoop := isInLoop(from, loop)
	cnt := 1
	for len(queue) > 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for _, v := range allNeighs(cur, grid) {
			if _, ok := visited[v]; !ok {
				visited[v] = true
				cnt += 1
				queue = append(queue, v)
				allInLoop = allInLoop && isInLoop(v, loop)
			}
		}
	}
	if allInLoop {
		return cnt
	}
	return 0
}

func main() {
	grid := readInput()
	loop := findLoop(grid)
	visited := make(map[Coord]bool)
	for _, v := range loop {
		visited[v] = true
	}
	part2 := 0
	for i, row := range grid {
		for j := range row {
			part2 += dfs(Coord{i, j}, grid, visited, loop)
		}
	}
	println("Part 1:", len(loop)/2)
	println("Part 2:", part2)
}
