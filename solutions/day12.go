package solutions

import (
	"AdventOfCode2024/utils"
	"math"
	"sort"
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
	count := 0
	for _, index := range region {
		neighbors := getNeighbours(index, w, h)
		count += 4 - len(neighbors)
		for j := 0; j < len(neighbors); j++ {
			n := neighbors[j]
			if part1Regions[n] != k {
				count++
			}
		}
	}
	return count
}

func getOrderedNeighbors(i, w, h int) []int {
	x, y := utils.OneDTwoD(i, w)
	neighbours := make([]int, 4)
	for j := 0; j < 4; j++ {
		neighbours[j] = -1
	}
	if x > 0 {
		neighbours[0] = utils.TwoDToOneD(x-1, y, w)
	}
	if x+1 < w {
		neighbours[1] = utils.TwoDToOneD(x+1, y, w)
	}
	if y > 0 {
		neighbours[2] = utils.TwoDToOneD(x, y-1, w)
	}
	if y+1 < h {
		neighbours[3] = utils.TwoDToOneD(x, y+1, w)
	}
	return neighbours
}

func getStraightSidePerimeter(k, w, h int) int {
	sides := []int{}
	region := part1RegionMap[k]
	for _, index := range region {
		neighbors := getOrderedNeighbors(index, w, h)
		for j := 0; j < len(neighbors); j++ {
			n := neighbors[j]
			hash := utils.SzudzikPairing(index+1, j+2)
			if n == -1 || part1Regions[n] != k {
				sides = append(sides, hash)
			}
		}
	}
	straightSides := make([]int, len(sides))
	for i := 0; i < len(sides); i++ {
		straightSides[i] = -1
	}

	for i := 0; i < len(sides); i++ {
		if straightSides[i] != -1 {
			continue
		}
		straightSides[i] = 1
		a := sides[i]
		maxXdist := 1
		maxYdist := 1
		for j := i + 1; j < len(sides); j++ {
			b := sides[j]
			ai, aj := utils.SzudzikUnpairing(a)
			ai, aj = ai-1, aj-2
			bi, bj := utils.SzudzikUnpairing(b)
			bi, bj = bi-1, bj-2
			ax, ay := utils.OneDTwoD(ai, w)
			bx, by := utils.OneDTwoD(bi, w)
			xdist := int(math.Abs(float64(ax - bx)))
			ydist := int(math.Abs(float64(ay - by)))
			if xdist == maxXdist && ydist == 0 {
				if aj == 2 && bj == 2 || aj == 3 && bj == 3 {
					straightSides[j] = 0
					maxXdist++
				}
				continue
			}
			if xdist == 0 && ydist == maxYdist {
				if aj == 0 && bj == 0 || aj == 1 && bj == 1 {
					straightSides[j] = 0
					maxYdist++
				}
				continue
			}
		}
	}
	count := 0
	for i := 0; i < len(straightSides); i++ {
		if straightSides[i] == 1 {
			count++
		}
	}
	return count
}

func solveDay12(parsed []int, w, h int) (int, int) {
	part1Regions = make([]int, len(parsed))
	for i := 0; i < len(parsed); i++ {
		part1Regions[i] = -1
	}
	part1RegionMap = map[int][]int{}
	for i := 0; i < len(parsed); i++ {
		expandRegion(parsed, i, i, w, h)
	}
	part1Price := 0
	part2Price := 0
	for k, v := range part1RegionMap {
		sort.Ints(v)
		area := len(v)
		part1Price += area * getPerimeter(k, w, h)
		part2Price += area * getStraightSidePerimeter(k, w, h)
	}
	return part1Price, part2Price
}

func Day12(input []string) []string {
	parsed := parseDay12(input)
	solution1, solution2 := solveDay12(parsed, len(input[0]), len(input))
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
