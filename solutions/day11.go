package solutions

import (
	"AdventOfCode2024/utils"
	"math"
	"strconv"
	"strings"
)

/*
If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
*/

func parseDay11(input []string) []int {
	line := input[0]
	numStrs := strings.Split(line, " ")
	nums := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		nums[i], _ = strconv.Atoi(numStr)
	}
	return nums
}

func intLength(num int) int {
	length := 0
	for num > 0 {
		num /= 10
		length++
	}
	return length
}

func splitStone(stone, length int) (int, int) {
	half := length / 2
	left := stone / int(math.Pow10(half))
	right := stone % int(math.Pow10(half))
	return left, right
}

var stoneCache = map[int]int{}

func recurseStoneCount(stone, depth int) int {
	if depth == 0 {
		return 1
	}
	hash := utils.SzudzikPairing(stone, depth)
	if val, ok := stoneCache[hash]; ok {
		return val
	}
	length := intLength(stone)
	result := 0
	if stone == 0 {
		result = recurseStoneCount(1, depth-1)
	} else if length%2 == 0 {
		left, right := splitStone(stone, length)
		result = recurseStoneCount(left, depth-1) + recurseStoneCount(right, depth-1)
	} else {
		result = recurseStoneCount(stone*2024, depth-1)
	}
	stoneCache[hash] = result
	return result
}

func solveDay11(parsed []int, n int) int {
	total := 0
	for _, stone := range parsed {
		total += recurseStoneCount(stone, n)
	}
	return total
}

func Day11(input []string) []string {
	stones := parseDay11(input)
	solution1 := solveDay11(stones, 25)
	solution2 := solveDay11(stones, 75)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
