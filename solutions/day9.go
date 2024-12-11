package solutions

import (
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
	result := 0
	rindex := 0
	maxUnhandled := len(parsed) - 1
	for i := 0; i < len(parsed); i++ {
		gap := i%2 == 1
		if !gap {
			if parsed[i] < 0 {
				rindex += (parsed[i] * -1)
				continue
			}
			val := i / 2
			for j := 0; j < parsed[i]; j++ {
				result += val * rindex
				rindex++
			}
			continue
		}
		gapSize := parsed[i]
		contiguous := true
		for e := maxUnhandled; e > i && gapSize > 0; e -= 2 {
			if parsed[e] < 0 {
				continue
			}
			if parsed[e] <= gapSize {
				for j := 0; j < parsed[e]; j++ {
					result += (e / 2) * rindex
					rindex++
				}
				gapSize -= parsed[e]
				parsed[e] *= -1
			} else {
				if contiguous {
					maxUnhandled = e
					contiguous = false
				}
			}
		}
		if gapSize > 0 {
			rindex += gapSize
		}
	}
	return result
}

func Day9(input []string) []string {
	parsed := parseDay9(input)
	solution1 := solveDay9Part1(parsed)
	solution2 := solveDay9Part2(parsed)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
