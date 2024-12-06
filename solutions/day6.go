package solutions

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strconv"
)

func parseDay6Input(input []string) ([]int, [][]bool) {
	currentPos := []int{0, 0}
	_map := [][]bool{}
	for i, line := range input {
		_map = append(_map, []bool{})
		for j, char := range line {
			if char == '^' {
				currentPos = []int{i, j}
				_map[i] = append(_map[i], false)
				continue
			}
			if char == '#' {
				_map[i] = append(_map[i], true)
				continue
			}
			_map[i] = append(_map[i], false)
		}
	}
	return currentPos, _map
}

func printMap(startPos []int, visited map[int][]bool, _map [][]bool, extraObstacle []int) {
	for i := 0; i < len(_map); i++ {
		for j := 0; j < len(_map[0]); j++ {
			hash := utils.TwoDToOneD(j, i, len(_map[0]))
			if extraObstacle != nil && i == extraObstacle[0] && j == extraObstacle[1] {
				fmt.Print("%")
				continue
			}
			if v, ok := visited[hash]; ok {
				if startPos[0] == i && startPos[1] == j {
					fmt.Print("O")
					continue
				}
				if v[0] || v[1] {
					if v[2] || v[3] {
						fmt.Print("+")
						continue
					}
					fmt.Print("-")
					continue
				}
				if v[2] || v[3] {
					fmt.Print("|")
				}
				continue
			}
			if _map[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func turnRight(currentDir []int) []int {
	if currentDir[0] == 0 {
		if currentDir[1] == 1 {
			return []int{1, 0}
		}
		return []int{-1, 0}
	}
	if currentDir[0] == 1 {
		return []int{0, -1}
	}
	return []int{0, 1}
}

func dirIndex(currentDir []int) int {
	if currentDir[0] == 0 {
		if currentDir[1] == 1 {
			return 0
		}
		return 1
	}
	if currentDir[0] == 1 {
		return 2
	}
	return 3
}

func solveDay6Part1(startPos []int, _map [][]bool, extraObstacle []int) (int, map[int][]bool) {
	visited := map[int][]bool{}
	currentDir := []int{-1, 0}
	currentPos := []int{startPos[0], startPos[1]}
	for true {
		if currentPos[0] < 0 || currentPos[0] >= len(_map) || currentPos[1] < 0 || currentPos[1] >= len(_map[0]) {
			break
		}
		if _map[currentPos[0]][currentPos[1]] ||
			(extraObstacle != nil && currentPos[0] == extraObstacle[0] && currentPos[1] == extraObstacle[1]) {
			currentPos[0] -= currentDir[0]
			currentPos[1] -= currentDir[1]
			currentDir = turnRight(currentDir)
			continue
		}
		hash := utils.TwoDToOneD(currentPos[1], currentPos[0], len(_map[0]))
		if v, ok := visited[hash]; ok {
			if v[dirIndex(currentDir)] {
				return -1, visited
			}
		} else {
			visited[hash] = []bool{false, false, false, false}
		}
		visited[hash][dirIndex(currentDir)] = true
		currentPos[0] += currentDir[0]
		currentPos[1] += currentDir[1]
	}
	return len(visited), visited
}

func solveDay6Part2(startPos []int, _map [][]bool, visited map[int][]bool) int {
	count := 0
	for k := range visited {
		x, y := utils.OneDTwoD(k, len(_map[0]))
		if x == startPos[1] && y == startPos[0] {
			continue
		}
		result, _ := solveDay6Part1(startPos, _map, []int{y, x})
		if result == -1 {
			count++
		}
	}
	return count
}

func Day6(input []string) []string {
	currentPos, _map := parseDay6Input(input)
	solution1, visited := solveDay6Part1(currentPos, _map, nil)
	solution2 := solveDay6Part2(currentPos, _map, visited)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
