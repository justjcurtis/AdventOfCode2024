package solutions

import (
	"AdventOfCode2024/utils"
	"strconv"
)

func parseDay12(input []string) []int {
	maxhash := utils.TwoDToOneD(len(input[0])-1, len(input)-1, len(input[0]))
	parsed := make([]int, len(input)*len(input[0]))
	fn := func(i int) {
		x, y := utils.OneDTwoD(i, len(input[0]))
		parsed[i] = int(input[y][x])
	}
	utils.ParalleliseVoid(fn, maxhash+1)
	return parsed
}

var part1RegionMap map[int][]int
var part1Regions []int

func expandRegion(parsed []int, o, i, w, h int) {
	if part1Regions[i] != -1 || parsed[i] != parsed[o] {
		return
	}
	part1Regions[i] = o
	part1RegionMap[o] = append(part1RegionMap[o], i)
	neighbors := getNeighbours(i, w, h)
	for _, n := range neighbors {
		expandRegion(parsed, o, n, w, h)
	}
}

func getPerimeter(k, w, h int) int {
	region := part1RegionMap[k]
	uniqueNeighbors := map[int]bool{}
	for index := range region {
		neighbors := getNeighbours(index, w, h)
		for _, n := range neighbors {
			if part1Regions[n] != k {
				uniqueNeighbors[part1Regions[n]] = true
			}
		}
	}
	return len(uniqueNeighbors)
}

func solveDay12Part1(parsed []int, w, h int) int {
	part1Regions = make([]int, len(parsed))
	for i := 0; i < len(parsed); i++ {
		part1Regions[i] = -1
	}
	part1RegionMap = map[int][]int{}
	for i := 0; i < len(parsed); i++ {
		expandRegion(parsed, i, i, w, h)
	}
	totalPrice := 0
	for k, v := range part1RegionMap {
		perimeter := getPerimeter(k, w, h)
		area := len(v)
		price := area * perimeter
		totalPrice += price
	}
	return totalPrice
}

func Day12(input []string) []string {
	parsed := parseDay12(input)
	solution1 := solveDay12Part1(parsed, len(input[0]), len(input))
	return []string{strconv.Itoa(solution1)}
}
