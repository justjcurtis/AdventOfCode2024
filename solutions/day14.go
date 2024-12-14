package solutions

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strconv"
	"strings"
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

func Day14(lines []string) []string {
	parsed := parseDay14(lines)
	fmt.Println(parsed)
	return []string{}
}
