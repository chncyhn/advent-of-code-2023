package main

import (
	"bufio"
	"os"
)

type Particle struct {
	position  [2]int
	direction [2]int
}

func move(p Particle) Particle {
	return Particle{[2]int{p.position[0] + p.direction[0], p.position[1] + p.direction[1]}, p.direction}
}

func reflectAngled(p Particle, device rune) Particle {
	pos := p.position
	if device == '\\' {
		return Particle{pos, [2]int{p.direction[1], p.direction[0]}}
	}
	if device == '/' {
		return Particle{pos, [2]int{-p.direction[1], -p.direction[0]}}
	}
	panic("unexpected device " + string(device))
}

func reflectStraight(p Particle) []Particle {
	pos := p.position
	if p.direction[0] == 0 {
		return []Particle{{pos, [2]int{1, 0}}, {pos, [2]int{-1, 0}}}
	} else {
		return []Particle{{pos, [2]int{0, 1}}, {pos, [2]int{0, -1}}}
	}
}

func isAligned(p Particle, device rune) bool {
	dir := p.direction
	return (dir[0] != 0 && device == '|') || (dir[1] != 0 && device == '-')

}

func simulateParticle(p Particle, grid [][]rune) []Particle {
	device := grid[p.position[0]][p.position[1]]
	if device == '.' || isAligned(p, device) {
		return []Particle{move(p)}
	}
	if device == '-' || device == '|' {
		return reflectStraight(p)
	}
	if device == '\\' || device == '/' {
		return []Particle{move(reflectAngled(p, device))}
	}
	panic("unexpected device " + string(device))
}

func inGrid(pos [2]int, grid [][]rune) bool {
	return pos[0] >= 0 && pos[0] < len(grid) && pos[1] >= 0 && pos[1] < len(grid[0])
}

func runSimulation(grid [][]rune, start Particle) int {
	currentParticles := []Particle{start}
	seenParticles := make(map[Particle]bool)
	seenParticles[start] = true
	size := 0
	for size != len(seenParticles) {
		size = len(seenParticles)
		nextParticles := []Particle{}
		for _, p := range currentParticles {
			for _, next := range simulateParticle(p, grid) {
				if !seenParticles[next] && inGrid(next.position, grid) {
					nextParticles = append(nextParticles, next)
				}
			}
		}
		for _, p := range nextParticles {
			seenParticles[p] = true
		}
		currentParticles = nextParticles
	}
	filledLocs := make(map[[2]int]bool)
	for p := range seenParticles {
		filledLocs[p.position] = true
	}
	return len(filledLocs)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	part1 := runSimulation(grid, Particle{[2]int{0, 0}, [2]int{0, 1}})
	println("Part 1:", part1)

	part2 := 0
	for i := 0; i < len(grid); i++ {
		part2 = max(part2, runSimulation(grid, Particle{[2]int{i, 0}, [2]int{0, 1}}))
		part2 = max(part2, runSimulation(grid, Particle{[2]int{i, len(grid[0]) - 1}, [2]int{0, -1}}))
	}
	for j := 0; j < len(grid[0]); j++ {
		part2 = max(part2, runSimulation(grid, Particle{[2]int{0, j}, [2]int{1, 0}}))
		part2 = max(part2, runSimulation(grid, Particle{[2]int{len(grid) - 1, j}, [2]int{-1, 0}}))
	}
	println("Part 2:", part2)
}
