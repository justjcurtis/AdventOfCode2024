package solutions

import (
	"AdventOfCode2024/utils"
	"sort"
	"strconv"
	"strings"
)

func parseDay1Line(line string) (int, int) {
	strs := strings.Split(line, "   ")
	l, _ := strconv.Atoi(strs[0])
	r, _ := strconv.Atoi(strs[1])
	return l, r
}

func parseDay1(input []string) ([]int, []int) {
	left := make([]int, len(input))
	right := make([]int, len(input))
	fn := func(i int) {
		l, r := parseDay1Line(input[i])
		left[i] = l
		right[i] = r
	}
	utils.ParalleliseVoid(fn, len(input))
	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

func solveDay1Part1(left []int, right []int) int {
	fn := func(i int) int {
		result := left[i] - right[i]
		if result < 0 {
			result = -result
		}
		return result
	}
	return utils.Parallelise(utils.IntAcc, fn, len(left))
}

func solveDay1Part2(left []int, right []int) int {
	leftMap := utils.CountMap(left)
	leftKeys := []int{}
	for k := range leftMap {
		leftKeys = append(leftKeys, k)
	}
	rightMap := utils.CountMap(right)
	fn := func(i int) int {
		l := leftKeys[i]
		if r, ok := rightMap[l]; ok {
			return l * r
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(leftKeys))
}

func Day1(input []string) []string {
	parsedLeft, parsedRight := parseDay1(input)
	solution1 := solveDay1Part1(parsedLeft, parsedRight)
	solution2 := solveDay1Part2(parsedLeft, parsedRight)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
