package solutions

import (
	"AdventOfCode2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

func parseDay14Line(line string) [][]int {
	split := strings.Split(line, " ")
	posStrs := strings.Split(split[0][2:], ",")
	x, _ := strconv.Atoi(posStrs[0])
	y, _ := strconv.Atoi(posStrs[1])
	pos := []int{x, y}
	velStrs := strings.Split(split[1][2:], ",")
	x, _ = strconv.Atoi(velStrs[0])
	y, _ = strconv.Atoi(velStrs[1])
	vel := []int{x, y}
	return [][]int{pos, vel}
}

func parseDay14(lines []string) [][][]int {
	parsed := make([][][]int, len(lines))
	fn := func(i int) {
		parsed[i] = parseDay14Line(lines[i])
	}
	utils.ParalleliseVoid(fn, len(lines))
	return parsed
}

func simulateDay14(pos, vel []int, w, h, seconds int) []int {
	newX := pos[0] + vel[0]*seconds
	newY := pos[1] + vel[1]*seconds
	for newX < 0 || newY < 0 {
		newX += w
		newY += h
	}
	newX %= w
	newY %= h
	return []int{newX, newY}
}

func solveDay14Part1(parsed [][][]int, w, h int) int {
	result := make([][]int, len(parsed))
	fn := func(i int) {
		pos := parsed[i][0]
		vel := parsed[i][1]
		result[i] = simulateDay14(pos, vel, w, h, 100)
	}
	utils.ParalleliseVoid(fn, len(parsed))
	midX, midY := w/2, h/2
	quadrants := make([]int, 4)
	for _, pos := range result {
		if pos[0] == midX || pos[1] == midY {
			continue
		}
		if pos[0] < midX && pos[1] < midY {
			quadrants[0]++
		} else if pos[0] >= midX && pos[1] < midY {
			quadrants[1]++
		} else if pos[0] < midX && pos[1] >= midY {
			quadrants[2]++
		} else {
			quadrants[3]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func calculateEntropyDay14(positions [][]int, w, h, chunkSize int) float64 {
	chunkCount := (w / chunkSize) * (h / chunkSize)
	chunkMap := make([]int, chunkCount)
	mu := sync.Mutex{}
	fn := func(i int) {
		pos := positions[i]
		chunkX := pos[0] / chunkSize
		chunkY := pos[1] / chunkSize
		hash := utils.TwoDToOneD(chunkX, chunkY, w/chunkSize)
		for hash >= chunkCount {
			if pos[0] > pos[1] {
				chunkX--
			} else {
				chunkY--
			}
			hash = utils.TwoDToOneD(chunkX, chunkY, w/chunkSize)
		}
		mu.Lock()
		chunkMap[hash]++
		mu.Unlock()
	}
	utils.ParalleliseVoid(fn, len(positions))

	entropy := 0.0
	fn = func(i int) {
		if chunkMap[i] == 0 {
			return
		}
		p := float64(chunkMap[i]) / float64(chunkCount)
		if p > 0 {
			mu.Lock()
			entropy -= p * math.Log2(p)
			mu.Unlock()
		}
	}
	utils.ParalleliseVoid(fn, chunkCount)

	return entropy
}

func printGridDay14(positions [][]int, w, h int) {
	positionsMap := make([]int, w*h)
	mu := sync.Mutex{}
	fn := func(i int) {
		pos := positions[i]
		hash := utils.TwoDToOneD(pos[0], pos[1], w)
		mu.Lock()
		positionsMap[hash]++
		mu.Unlock()
	}
	utils.ParalleliseVoid(fn, len(positions))
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			hash := utils.TwoDToOneD(j, i, w)
			if positionsMap[hash] == 0 {
				fmt.Print(".")
			} else if positionsMap[hash] < 10 {
				fmt.Print(positionsMap[hash])
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func solveDay14Part2(parsed [][][]int, w, h int) int {
	result := make([][]int, len(parsed))
	index := 0
	entropy := math.MaxFloat64
	tree := [][]int{}
	for n := 0; n < 7441; n++ {
		fn := func(i int) {
			pos := parsed[i][0]
			vel := parsed[i][1]
			result[i] = simulateDay14(pos, vel, w, h, n)
		}
		utils.ParalleliseVoid(fn, len(parsed))
		newEntropy := calculateEntropyDay14(result, w, h, 10)
		if newEntropy < entropy {
			entropy = newEntropy
			index = n
			tree = make([][]int, len(result))
			for i := 0; i < len(result); i++ {
				tree[i] = make([]int, 2)
				copy(tree[i], result[i])
			}
		}
	}
	// printGridDay14(tree, w, h)
	return index
}

func Day14(lines []string) []string {
	parsed := parseDay14(lines)
	w, h := 101, 103
	if len(parsed) == 12 {
		w, h = 11, 7
	}
	solution1 := solveDay14Part1(parsed, w, h)
	solution2 := 0
	if len(parsed) > 12 {
		solution2 = solveDay14Part2(parsed, w, h)
	}
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
