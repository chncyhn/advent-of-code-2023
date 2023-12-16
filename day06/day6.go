package main

import (
	"bufio"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func mapToInt(vals []string) (ret []int) {
	for _, s := range vals {
		dig, _ := strconv.Atoi(s)
		ret = append(ret, dig)
	}
	return
}

func quadraticRoots(a, b, c int) (float64, float64) {
	disc := float64(b*b - 4*a*c)
	return (-float64(b) + math.Sqrt(disc)) / (2 * float64(a)), (-float64(b) - math.Sqrt(disc)) / (2 * float64(a))
}

func roundDown(f float64) int {
	return int(math.Floor(f))
}

func roundUp(f float64) int {
	return int(math.Ceil(f))
}

func solve(t int, d int) int {
	a, b := quadraticRoots(-1, t, -d)
	lo, hi := roundUp(a+1e-8), roundDown(b-1e-8)
	return hi - lo + 1
}

func part1(times []int, distances []int) {
	ret := 1
	for i, t := range times {
		ret *= solve(t, distances[i])
	}
	println(ret)
}

func quadraticRootsBig(a, b, c *big.Int) (*big.Float, *big.Float) {
	disc := new(big.Int).Mul(b, b)
	disc.Sub(disc, new(big.Int).Mul(big.NewInt(4), new(big.Int).Mul(a, c)))
	discSqrt := new(big.Float).SetPrec(256)
	discSqrt.Sqrt(new(big.Float).SetInt(disc))
	negB := new(big.Float).Neg(new(big.Float).SetInt(b))
	twoA := new(big.Float).Mul(big.NewFloat(2), new(big.Float).SetInt(a))
	root1 := new(big.Float).Quo(new(big.Float).Add(negB, discSqrt), twoA)
	root2 := new(big.Float).Quo(new(big.Float).Sub(negB, discSqrt), twoA)
	return root1, root2
}

func part2(times []int, distances []int) {
	time := big.NewInt(0)
	dist := big.NewInt(0)
	for i, t := range times {
		time.Mul(time, big.NewInt(int64(math.Pow(10, math.Ceil(math.Log10(float64(t)))))))
		dist.Mul(dist, big.NewInt(int64(math.Pow(10, math.Ceil(math.Log10(float64(distances[i])))))))
		time.Add(time, big.NewInt(int64(t)))
		dist.Add(dist, big.NewInt(int64(distances[i])))
	}
	a, b := quadraticRootsBig(big.NewInt(-1), time, new(big.Int).Neg(dist))
	af, _ := a.Float64()
	bf, _ := b.Float64()
	lo, hi := roundUp(af+1e-8), roundDown(bf-1e-8)
	println(hi - lo + 1)
}

func main() {
	file, _ := os.Open("full.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	times := mapToInt(strings.Fields(scanner.Text())[1:])
	scanner.Scan()
	distances := mapToInt(strings.Fields(scanner.Text())[1:])

	part1(times, distances)
	part2(times, distances)
}
