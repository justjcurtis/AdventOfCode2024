package solutions

import (
	"strconv"
	"sync"
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
				val := endex / 2
				n := unused
				result += val * (n * (2*rindex + (n - 1)) / 2)
				rindex += n
			}
			break
		}
		aIndex := index / 2
		aAmnt := parsed[index]
		result += aIndex * (aAmnt * (2*rindex + (aAmnt - 1)) / 2)
		rindex += aAmnt
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
	toHandle := make([]int, (len(parsed)/2)+1)
	for i := 0; i < len(parsed); i += 2 {
		toHandle[i/2] = i
	}
	for i := 0; i < len(parsed); i++ {
		gap := i%2 == 1
		if !gap {
			n := parsed[i]
			if n < 0 {
				rindex += (n * -1)
				continue
			}
			val := i / 2
			result += val * (n * (2*rindex + (n - 1)) / 2)
			rindex += n
			continue
		}
		gapSize := parsed[i]
		for e := len(toHandle) - 1; toHandle[e] > i && gapSize > 0; e-- {
			endex := toHandle[e]
			if parsed[endex] < 0 {
				continue
			}
			if parsed[endex] <= gapSize {
				val := endex / 2
				n := parsed[endex]
				result += val * (n * (2*rindex + (n - 1)) / 2)
				rindex += n
				gapSize -= n
				parsed[endex] *= -1
				toHandle = append(toHandle[:e], toHandle[e+1:]...)
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
	wg := sync.WaitGroup{}
	var solution1 int
	var solution2 int
	wg.Add(2)
	go func() {
		defer wg.Done()
		solution1 = solveDay9Part1(parsed)
	}()
	go func() {
		defer wg.Done()
		solution2 = solveDay9Part2(parsed)
	}()
	wg.Wait()
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
