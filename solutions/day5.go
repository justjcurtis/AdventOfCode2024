package solutions

import (
	"AdventOfCode2024/utils"
	"strconv"
	"strings"
	"sync"
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
	fn := func() {
		for _, line := range input {
			if line == "" {
				break
			}
			a, b := parseRule(line)
			rules.addRule(a, b)
		}
	}
	updatesList := [][]int{}
	fn2 := func() {
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
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fn()
	}()
	go func() {
		defer wg.Done()
		fn2()
	}()
	wg.Wait()

	return rules, updatesList
}

func checkUpdate(i int, updates []int, rules day5Rules) int {
	a := updates[i]
	for j, b := range updates[i+1:] {
		if _, ok := rules.afters[b]; ok {
			if _, ok := rules.afters[b][a]; ok {
				return j + i + 1
			}
		}
	}
	return -1
}

func checkUpdates(updates []int, rules day5Rules, start int) (int, int) {
	for i := start; i < len(updates)-1; i++ {
		j := checkUpdate(i, updates, rules)
		if j > -1 {
			return i, j
		}
	}
	return -1, -1
}

func solveDay5Part1(rules day5Rules, updatesList [][]int) int {
	fn := func(index int) int {
		updates := updatesList[index]
		if faili, _ := checkUpdates(updates, rules, 0); faili > -1 {
			return 0
		}
		return updates[(len(updates) / 2)]
	}
	result := utils.Parallelise(utils.IntAcc, fn, len(updatesList))
	return result
}

func swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func solveDay5Part2(rules day5Rules, updatesList [][]int) int {
	fn := func(index int) int {
		updates := updatesList[index]
		faili, failj := checkUpdates(updates, rules, 0)
		if faili == -1 {
			return 0
		}
		swap(faili, failj, updates)
		for faili > -1 {
			newFaili, newFailj := checkUpdates(updates, rules, faili)
			if newFaili == -1 {
				break
			}
			swap(newFaili, newFailj, updates)
			faili, failj = newFaili, newFailj
		}
		return updates[(len(updates) / 2)]
	}
	return utils.Parallelise(utils.IntAcc, fn, len(updatesList))
}

func Day5(input []string) []string {
	rules, updatesList := parseDay5Input(input)
	var solution1, solution2 int
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		solution1 = solveDay5Part1(rules, updatesList)
	}()
	go func() {
		defer wg.Done()
		updatesCopy := make([][]int, len(updatesList))
		for i, updates := range updatesList {
			updatesCopy[i] = make([]int, len(updates))
			copy(updatesCopy[i], updates)
		}
		solution2 = solveDay5Part2(rules, updatesCopy)
	}()
	wg.Wait()
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
