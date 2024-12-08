package solutions

import (
	"AdventOfCode2024/utils"
	"math"
	"strconv"
	"strings"
	"sync"
)

type day7Op struct {
	id string
	op func(a, b int) int
}

var day7Part1Ops = []day7Op{
	{"+", func(a, b int) int { return a + b }},
	{"*", func(a, b int) int { return a * b }},
}

func parseDay7Line(line string) []int {
	var buffer strings.Builder
	index := 0
	for i, char := range line {
		if char == ':' {
			index = i
			break
		}
		buffer.WriteRune(char)
	}
	testValue, _ := strconv.Atoi(buffer.String())
	nums := []int{testValue}
	buffer.Reset()
	for i := index + 2; i < len(line); i++ {
		if line[i] == ' ' {
			num, _ := strconv.Atoi(buffer.String())
			nums = append(nums, num)
			buffer.Reset()
			continue
		}
		buffer.WriteByte(line[i])
	}
	num, _ := strconv.Atoi(buffer.String())
	nums = append(nums, num)
	return nums
}

func parseDay7(input []string) [][]int {
	parsed := make([][]int, len(input))
	fn := func(i int) {
		parsed[i] = parseDay7Line(input[i])
	}
	utils.ParalleliseVoid(fn, len(input))
	return parsed

}

func calculateResult(line []int, ops []day7Op, currentOps []int) bool {
	value := line[1]
	for i := 2; i < len(line); i++ {
		value = ops[currentOps[i-2]].op(value, line[i])
	}
	result := value == line[0]
	return result
}

func recurseDay7(line []int, ops []day7Op, currentOps []int, depth int) bool {
	if depth == len(line)-2 {
		return calculateResult(line, ops, currentOps)
	}
	for i := range ops {
		currentOps[depth] = i
		if recurseDay7(line, ops, currentOps, depth+1) {
			return true
		}
	}
	return false
}

var day7SkipCache = sync.Map{}

func solveDay7Part1(parsed [][]int) int {
	fn := func(i int) int {
		line := parsed[i]
		currentOps := make([]int, len(line)-2)
		if recurseDay7(line, day7Part1Ops, currentOps, 0) {
			day7SkipCache.Store(i, true)
			return line[0]
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(parsed))
}

func getIntLength(num int) int {
	switch {
	case num < 10:
		return 1
	case num < 100:
		return 2
	case num < 1000:
		return 3
	default:
		length := 0
		for num > 0 {
			num /= 10
			length++
		}
		return length
	}
}

var day7Part2Ops = []day7Op{
	day7Part1Ops[0],
	day7Part1Ops[1],
	{"||", func(a, b int) int {
		blength := getIntLength(b)
		result := (a * int(math.Pow10(blength))) + b
		return result
	}},
}

func solveDay7Part2(parsed [][]int) int {
	fn := func(i int) int {
		if _, ok := day7SkipCache.Load(i); ok {
			return 0
		}
		line := parsed[i]
		currentOps := make([]int, len(line)-2)
		if recurseDay7(line, day7Part2Ops, currentOps, 0) {
			return line[0]
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(parsed))
}

func Day7(input []string) []string {
	parsed := parseDay7(input)
	solution1 := solveDay7Part1(parsed)
	solution2 := solveDay7Part2(parsed) + solution1
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
