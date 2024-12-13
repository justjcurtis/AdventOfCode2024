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

func minTokenCost(ax, ay, bx, by, px, py, ac, bc int, limit bool) int {
	minCost := -1
	for n := 0; n <= 100 || !limit; n++ {
		m := float64((n*bx)-px) / float64(-ax)
		if m < 0 {
			break
		}
		intM := int(m)
		if m == float64(intM) {
			if intM*ay+n*by == py {
				cost := n*bc + intM*ac
				if minCost == -1 || cost < minCost {
					minCost = cost
				}
			}
		}
	}
	if minCost != -1 {
		return minCost
	}
	return 0
}

func solveDay13Part1(parsed [][][]int) int {
	fn := func(i int) int {
		entry := parsed[i]
		ax, ay := entry[0][0], entry[0][1]
		bx, by := entry[1][0], entry[1][1]
		px, py := entry[2][0], entry[2][1]
		cost := minTokenCost(ax, ay, bx, by, px, py, 3, 1, true)
		return cost
	}
	return utils.Parallelise(utils.IntAcc, fn, len(parsed))
}

func Day13(input []string) []string {
	parsed := parseDay13(input)
	solution1 := solveDay13Part1(parsed)
	return []string{strconv.Itoa(solution1)}
}
