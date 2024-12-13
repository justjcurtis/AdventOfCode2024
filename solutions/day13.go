package solutions

import (
	"AdventOfCode2024/utils"
	"strconv"
	"strings"
)

func parseDay13(input []string) [][][]int {
	parsed := [][][]int{{}}
	for _, line := range input {
		if line == "" {
			parsed = append(parsed, [][]int{})
			continue
		}
		var numstrs []string
		if line[0] == 'B' {
			line = line[12:]
			numstrs = strings.Split(line, ", Y+")
		} else {
			line = line[9:]
			numstrs = strings.Split(line, ", Y=")
		}
		x, _ := strconv.Atoi(numstrs[0])
		y, _ := strconv.Atoi(numstrs[1])
		nums := []int{x, y}
		parsed[len(parsed)-1] = append(parsed[len(parsed)-1], nums)
	}
	return parsed
}

func minTokenCost(ax, ay, bx, by, px, py, ac, bc int) int {
	diff := float64(ax*by - ay*bx)
	if diff == 0 {
		return 0
	}
	m := float64(px*by-py*bx) / diff
	n := float64(py*ax-px*ay) / diff
	if m != float64(int(m)) || n != float64(int(n)) {
		return 0
	}
	intm := int(m)
	intn := int(n)
	return intm*ac + intn*bc
}

func solveDay13(parsed [][][]int, offset int) int {
	fn := func(i int) int {
		entry := parsed[i]
		ax, ay := entry[0][0], entry[0][1]
		bx, by := entry[1][0], entry[1][1]
		px, py := entry[2][0], entry[2][1]
		cost := minTokenCost(ax, ay, bx, by, px+offset, py+offset, 3, 1)
		return cost
	}
	return utils.Parallelise(utils.IntAcc, fn, len(parsed))
}

func Day13(input []string) []string {
	parsed := parseDay13(input)
	solution1 := solveDay13(parsed, 0)
	solution2 := solveDay13(parsed, 10000000000000)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
