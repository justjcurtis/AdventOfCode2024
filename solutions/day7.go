package solutions

import (
	"AdventOfCode2024/utils"
	"math"
	"strconv"
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
	buffer := ""
	index := 0
	for i, char := range line {
		if char == ':' {
			index = i
			break
		}
		buffer += string(char)
	}
	testValue, _ := strconv.Atoi(buffer)
	nums := []int{testValue}
	buffer = ""
	for i := index + 2; i < len(line); i++ {
		if line[i] == ' ' {
			num, _ := strconv.Atoi(buffer)
			nums = append(nums, num)
			buffer = ""
			continue
		}
		buffer += string(line[i])
	}
	num, _ := strconv.Atoi(buffer)
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

func calculateResult(line []int, ops []day7Op, currentOps []int) int {
	target := line[0]
	value := line[1]
	for i := 2; i < len(line); i++ {
		value = ops[currentOps[i-2]].op(value, line[i])
	}
	if value == target {
		return value
	}
	return 0
}

func recurseDay7(line []int, ops []day7Op, currentOps []int) int {
	if len(currentOps) == len(line)-2 {
		return calculateResult(line, ops, currentOps)
	}
	total := 0
	for i := range ops {
		newOps := []int{}
		newOps = append(newOps, currentOps...)
		newOps = append(newOps, i)
		total += recurseDay7(line, ops, newOps)
		if total > 0 {
			break
		}
	}
	return total
}

func solveDay7Part1(parsed [][]int) int {
	fn := func(i int) int {
		return recurseDay7(parsed[i], day7Part1Ops, []int{})
	}
	return utils.Parallelise(utils.IntAcc, fn, len(parsed))
}

func getIntLength(num int) int {
	length := 0
	use := num
	for use > 0 {
		use /= 10
		length++
	}
	return length
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
		return recurseDay7(parsed[i], day7Part2Ops, []int{})
	}
	return utils.Parallelise(utils.IntAcc, fn, len(parsed))
}

func Day7(input []string) []string {
	parsed := parseDay7(input)
	solution1 := solveDay7Part1(parsed)
	solution2 := solveDay7Part2(parsed)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
