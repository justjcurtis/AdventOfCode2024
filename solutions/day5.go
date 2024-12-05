package solutions

import (
	"AdventOfCode2024/utils"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type day5Rules struct {
	afters map[int]map[int]bool
}

func (d *day5Rules) addRule(a, b int) {
	if _, ok := d.afters[a]; !ok {
		d.afters[a] = map[int]bool{}
	}
	d.afters[a][b] = true
}

func parseRule(line string) (int, int) {
	split := strings.Split(line, "|")
	a, _ := strconv.Atoi(split[0])
	b, _ := strconv.Atoi(split[1])
	return a, b
}

func parseUpdate(line string) []int {
	split := strings.Split(line, ",")
	update := []int{}
	for _, s := range split {
		a, _ := strconv.Atoi(s)
		update = append(update, a)
	}
	return update
}

func parseDay5Input(input []string) (day5Rules, [][]int) {
	rules := day5Rules{afters: map[int]map[int]bool{}}
	updatesList := [][]int{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, line := range input {
			if line == "" {
				break
			}
			a, b := parseRule(line)
			rules.addRule(a, b)
		}
	}()
	go func() {
		defer wg.Done()
		foundBreak := false
		for _, line := range input {
			if !foundBreak {
				if line == "" {
					foundBreak = true
					continue
				}
				continue
			}
			updatesList = append(updatesList, parseUpdate(line))
		}
	}()
	wg.Wait()

	return rules, updatesList
}

func checkUpdatesAndGetMiddle(rules day5Rules, updates []int) int {
	lastCount := len(updates)
	middleCount := len(updates) / 2
	middleIndex := -1
	failed := false
	for i := 0; i < len(updates); i++ {
		count := 0
		if v, ok := rules.afters[updates[i]]; ok {
			for j := 0; j < len(updates); j++ {
				if i == j {
					continue
				}
				if _, ok := v[updates[j]]; ok {
					count++
				}
			}
		}
		if middleIndex == -1 && count == middleCount {
			middleIndex = i
		}
		if count > lastCount {
			failed = true
		}
		if failed && middleIndex != -1 {
			return updates[middleIndex]
		}
		lastCount = count
	}
	return -1
}

func solveDay5Part1And2(rules day5Rules, updatesList [][]int) (int, int) {
	part1 := int32(0)
	part2 := int32(0)
	fn := func(index int) {
		updates := updatesList[index]
		if middle := checkUpdatesAndGetMiddle(rules, updates); middle != -1 {
			atomic.AddInt32(&part2, int32(middle))
			return
		}
		atomic.AddInt32(&part1, int32(updates[(len(updates)/2)]))

	}
	utils.ParalleliseVoid(fn, len(updatesList))
	return int(part1), int(part2)
}

func Day5(input []string) []string {
	rules, updatesList := parseDay5Input(input)
	solution1, solution2 := solveDay5Part1And2(rules, updatesList)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
