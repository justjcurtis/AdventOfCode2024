package solutions

import (
	"fmt"
	"strconv"
)

func parseDay9(input []string) []int {
	nums := []int{}
	for i := 0; i < len(input); i++ {
		line := input[i]
		for j := 0; j < len(line); j++ {
			nums = append(nums, int(line[j])-48)
		}
	}
	return nums
}

func solveDay9Part1(parsed []int) int {
	index := 0
	endex := len(parsed) - 1
	result := 0
	rindex := 0
	unused := 0
	for true {
		if index >= endex {
			if unused > 0 {
				for i := 0; i < unused; i++ {
					result += (endex / 2) * rindex
					rindex++
				}
			}
			break
		}
		aIndex := index / 2
		aAmnt := parsed[index]
		for i := 0; i < aAmnt; i++ {
			result += aIndex * rindex
			rindex++
		}
		gaps := parsed[index+1]
		index += 2
		bIndex := endex / 2
		if unused == 0 {
			unused = parsed[endex]
		}
		for i := 0; i < gaps; i++ {
			if unused == 0 {
				endex -= 2
				if endex <= index {
					break
				}
				bIndex = endex / 2
				unused = parsed[endex]
			}
			result += bIndex * rindex
			rindex++
			unused--
		}
		if unused == 0 {
			endex -= 2
		}
	}
	return result
}

func solveDay9Part2(parsed []int) int {
	// TODO: Implement part 2
	// I think I can do this without building processed
	// I can just keep track of the last gap index that was used up & how much of it remains if any
	// then I should be able to use a similar solution to part 1
	return 0
}

func Day9(input []string) []string {
	parsed := parseDay9(input)
	solution1 := solveDay9Part1(parsed)
	return []string{strconv.Itoa(solution1)}
}
