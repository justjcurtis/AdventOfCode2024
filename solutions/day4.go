package solutions

import (
	"AdventOfCode2024/utils"
	"strconv"
	"strings"
)

var noDirection = []int{0, 0}
var up = []int{0, -1}
var down = []int{0, 1}
var left = []int{-1, 0}
var right = []int{1, 0}
var upLeft = []int{-1, -1}
var upRight = []int{1, -1}
var downLeft = []int{-1, 1}
var downRight = []int{1, 1}
var directions = [][]int{up, down, left, right, upLeft, upRight, downLeft, downRight}

func getNeighbors(x, y, n int, d []int, input []string) []string {
	if n == 0 {
		return []string{}
	}
	if x < 0 || x >= len(input[0]) || y < 0 || y >= len(input) {
		return []string{}
	}
	if d[0] == 0 && d[1] == 0 {
		neighbors := []string{}
		for _, dir := range directions {
			next := append([]string{string(input[y][x])}, getNeighbors(x+dir[0], y+dir[1], n-1, dir, input)...)
			neighbors = append(neighbors, strings.Join(next, ""))
		}
		filtered := []string{}
		for _, neighbor := range neighbors {
			if len(neighbor) == n {
				filtered = append(filtered, neighbor)
			}
		}
		return filtered
	}
	return append([]string{string(input[y][x])}, getNeighbors(x+d[0], y+d[1], n-1, d, input)...)
}

func solveDay4Part1(input []string) int {
	fn := func(z int) int {
		xmasCount := 0
		x, y := utils.OneDTwoD(z, len(input[0]))
		char := input[y][x]
		if char == 'X' {
			neighbors := getNeighbors(x, y, 4, noDirection, input)
			for _, neighbor := range neighbors {
				if neighbor == "XMAS" {
					xmasCount++
				}
			}
		}
		return xmasCount
	}
	return utils.Parallelise(utils.IntAcc, fn, len(input)*len(input[0]))
}

func getCross(x, y int, input []string) []string {
	cross := []byte{}
	offsets := [][]int{upLeft, upRight, downLeft, downRight}
	for _, offset := range offsets {
		X := x + offset[0]
		Y := y + offset[1]
		if X < 0 || X >= len(input[0]) || Y < 0 || Y >= len(input) {
			return []string{}
		}
		char := input[Y][X]
		if char != 'M' && char != 'S' {
			return []string{}
		}
		cross = append(cross, char)
	}
	curr := input[y][x]
	strs := []string{string(cross[0] + curr + cross[3]), string(cross[1] + curr + cross[2])}
	return strs
}

func solveDay4Part2(input []string) int {
	fn := func(z int) int {
		x, y := utils.OneDTwoD(z, len(input[0]))
		char := input[y][x]
		if char == 'A' {
			crosses := getCross(x, y, input)
			if len(crosses) == 0 {
				return 0
			}
			if crosses[0] == "á" && crosses[1] == "á" {
				return 1
			}
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
