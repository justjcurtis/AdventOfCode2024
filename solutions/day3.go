package solutions

import (
	"AdventOfCode2024/utils"
	"regexp"
	"strconv"
	"strings"
)

func getIntsFromMulString(s string) (int, int) {
	reg := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := reg.FindStringSubmatch(s)
	a, _ := strconv.Atoi(matches[1])
	b, _ := strconv.Atoi(matches[2])
	return a, b
}

func solveDay3Part1(input []string) int {
	reg := regexp.MustCompile(`mul\(\d+,\d+\)`)
	fn := func(i int) int {
		line := input[i]
		matches := reg.FindAllString(line, -1)
		result := 0
		for _, match := range matches {
			a, b := getIntsFromMulString(match)
			result += a * b
		}
		return result
	}
	return utils.Parallelise(utils.IntAcc, fn, len(input))
}

func getOnlyDos(input []string) []string {
	reg := regexp.MustCompile(`don't\(\)(.*?)do\(\)`)
	joined := strings.Join(input, "") + "do()"
	return []string{reg.ReplaceAllString(joined, "")}
}

func solveDay3Part2(input []string) int {
	onlyDos := getOnlyDos(input)
	return solveDay3Part1(onlyDos)
}

func Day3(input []string) []string {
	solution1 := solveDay3Part1(input)
	solution2 := solveDay3Part2(input)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
