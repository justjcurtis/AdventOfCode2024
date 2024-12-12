package solutions

import (
	"AdventOfCode2024/utils"
	"strconv"
	"sync"
)

var day10Part1Cache sync.Map

func parseDay10(input []string) []int {
	maxhash := utils.TwoDToOneD(len(input[0])-1, len(input)-1, len(input[0]))
	parsed := make([]int, len(input)*len(input[0]))
	fn := func(i int) {
		x, y := utils.OneDTwoD(i, len(input[0]))
		parsed[i] = int(input[y][x]) - 48
	}
	utils.ParalleliseVoid(fn, maxhash+1)
	return parsed
}

func getNeighbours(i, w, h int) []int {
	x, y := utils.OneDTwoD(i, w)
	neighbours := []int{}
	if x > 0 {
		neighbours = append(neighbours, utils.TwoDToOneD(x-1, y, w))
	}
	if x+1 < w {
		neighbours = append(neighbours, utils.TwoDToOneD(x+1, y, w))
	}
	if y > 0 {
		neighbours = append(neighbours, utils.TwoDToOneD(x, y-1, w))
	}
	if y+1 < h {
		neighbours = append(neighbours, utils.TwoDToOneD(x, y+1, w))
	}
	return neighbours
}

func getTrailheadScore(parsed []int, o, i, w, h, current int) {
	if parsed[i] == 9 {
		hash := utils.SzudzikPairing(o, i)
		currentValue, ok := day10Part1Cache.Load(hash)
		if ok {
			day10Part1Cache.Store(hash, currentValue.(int)+1)
		} else {
			day10Part1Cache.Store(hash, 1)
		}
		return
	}
	neighbours := getNeighbours(i, w, h)
	for _, n := range neighbours {
		if parsed[n] == current {
			getTrailheadScore(parsed, o, n, w, h, current+1)
		}
	}
}

func solveDay10(parsed []int, w, h int) (int, int) {
	day10Part1Cache = sync.Map{}
	fn := func(i int) {
		if parsed[i] != 0 {
			return
		}
		getTrailheadScore(parsed, i, i, w, h, 1)
	}
	utils.ParalleliseVoid(fn, len(parsed))
	distinct := 0
	day10Part1Cache.Range(func(_, count interface{}) bool {
		distinct += count.(int)
		return true
	})

	return utils.SyncMapLength(&day10Part1Cache), int(distinct)
}

func Day10(input []string) []string {
	parsed := parseDay10(input)
	solution1, solution2 := solveDay10(parsed, len(input[0]), len(input))
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
