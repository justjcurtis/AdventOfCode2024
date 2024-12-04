package solutions

import (
	"AdventOfCode2024/utils"
	"math"
	"strconv"
)

var up = []int{0, -1}
var down = []int{0, 1}
var left = []int{-1, 0}
var right = []int{1, 0}
var upLeft = []int{-1, -1}
var upRight = []int{1, -1}
var downLeft = []int{-1, 1}
var downRight = []int{1, 1}
var directions = [][]int{up, down, left, right, upLeft, upRight, downLeft, downRight}

func getNeighborsRunner(x, y, n int, d []int, input []string) byte {
	if n == 0 {
		return 0
	}
	if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
		return 0
	}
	curr := input[y][x] * byte(math.Pow(2, float64(n-1)))
	return curr + getNeighborsRunner(x+d[0], y+d[1], n-1, d, input)
}

func getNeighbors(x, y, n int, input []string) int {
	xmasCount := 0
	curr := input[y][x] * byte(math.Pow(2, float64(n-1)))
	for _, dir := range directions {
		if int(curr+getNeighborsRunner(x+dir[0], y+dir[1], n-1, dir, input)) == 201 {
			xmasCount++
		}
	}
	return xmasCount
}

func solveDay4Part1(input []string) int {
	fn := func(z int) int {
		x, y := utils.OneDTwoD(z, len(input[0]))
		char := input[y][x]
		if char == 'X' {
			return getNeighbors(x, y, 4, input)
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(input)*len(input[0]))
}

func getByte(coords [][]int, input []string) byte {
	result := byte(0)
	for _, coord := range coords {
		x, y := coord[0], coord[1]
		result += input[y][x]
	}
	return result
}

func addCoords(x, y int, b []int) []int {
	return []int{x + b[0], y + b[1]}
}

func getCross(x, y int, input []string) int {
	if x < 1 || x >= len(input[0])-1 || y < 1 || y >= len(input)-1 {
		return 0
	}
	curr := input[y][x]
	aoffsets := [][]int{addCoords(x, y, upLeft), addCoords(x, y, downRight)}
	a := int(curr + getByte(aoffsets, input))
	boffsets := [][]int{addCoords(x, y, upRight), addCoords(x, y, downLeft)}
	b := int(curr + getByte(boffsets, input))
	if a == 225 && b == 225 {
		return 1
	}
	return 0
}

func solveDay4Part2(input []string) int {
	fn := func(z int) int {
		x, y := utils.OneDTwoD(z, len(input[0]))
		char := input[y][x]
		if char == 'A' {
			return getCross(x, y, input)
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(input)*len(input[0]))
}

func Day4(input []string) []string {
	solution1 := solveDay4Part1(input)
	solution2 := solveDay4Part2(input)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
