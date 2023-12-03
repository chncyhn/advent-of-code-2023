package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	red, blue, green int
}

type Game struct {
	rounds []Round
}

func part1(games []Game) {
	total := 0
	for i, game := range games {
		if feasibleGame(game) {
			total += (i + 1)
		}
	}
	fmt.Println(total)
}

func feasibleGame(game Game) bool {
	for _, round := range game.rounds {
		if !feasibleRound(round) {
			return false
		}
	}
	return true
}

func feasibleRound(round Round) bool {
	return round.red <= 12 && round.blue <= 14 && round.green <= 13
}

func part2(games []Game) {
	total := 0
	for _, game := range games {
		var maxBlue, maxRed, maxGreen int
		for _, round := range game.rounds {
			if round.red > maxRed {
				maxRed = round.red
			}
			if round.blue > maxBlue {
				maxBlue = round.blue
			}
			if round.green > maxGreen {
				maxGreen = round.green
			}
		}
		total += maxRed * maxBlue * maxGreen
	}
	fmt.Println(total)
}

func readGames() (games []Game) {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var roundsStrs = strings.Split(strings.Split(scanner.Text(), ": ")[1], "; ")
		var game Game
		for _, roundStr := range roundsStrs {
			var round Round
			for _, ballStr := range strings.Split(roundStr, ", ") {
				var ball = strings.Split(ballStr, " ")
				var val, _ = strconv.Atoi(ball[0])
				switch ball[1] {
				case "red":
					round.red = val
				case "blue":
					round.blue = val
				case "green":
					round.green = val
				}
			}
			game.rounds = append(game.rounds, round)
		}
		games = append(games, game)
	}
	return
}

func main() {
	games := readGames()
	part1(games)
	part2(games)
}
