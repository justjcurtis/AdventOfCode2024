package solutions

import (
	"fmt"
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

func solveDay11Part1(stones []int) int {
	current := make([]int, len(stones))
	copy(current, stones)
	for n := 0; n < 25; n++ {
		next := []int{}
		for _, stone := range current {
			length := intLength(stone)
			if stone == 0 {
				next = append(next, 1)
			} else if length%2 == 0 {
				left, right := splitStone(stone, length)
				next = append(next, left, right)
			} else {
				next = append(next, stone*2024)
			}
		}
		current = next
	}
	return len(current)
}

func experiment(parsed []int) {
	seen := make(map[int]bool)
	current := make([]int, len(parsed))
	copy(current, parsed)
	n := 0
	for true {
		next := []int{}
		noChange := true
		for _, stone := range current {
			if _, ok := seen[stone]; ok {
				continue
			}
			noChange = false
			seen[stone] = true
			length := intLength(stone)
			if stone == 0 {
				next = append(next, 1)
			} else if length%2 == 0 {
				left, right := splitStone(stone, length)
				next = append(next, left, right)
			} else {
				next = append(next, stone*2024)
			}
		}
		if noChange {
			break
		}
		current = next
		n++
	}
	fmt.Println("seeen all stones after", n, "iterations")
}

func Day11(input []string) []string {
	stones := parseDay11(input)
	experiment(stones)
	solution1 := solveDay11Part1(stones)
	return []string{strconv.Itoa(solution1)}
}
