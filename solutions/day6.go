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

func printMap(startPos []int, visited map[int][]int, _map [][]bool, extraObstacle []int) {
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
				if v[1] >= 0 || v[3] >= 0 {
					if v[0] >= 0 || v[2] >= 0 {
						fmt.Print("+")
						continue
					}
					fmt.Print("-")
					continue
				}
				if v[0] >= 0 || v[2] >= 0 {
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

var day6Directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func turnRight(dirIndex int) ([]int, int) {
	dirIndex = (dirIndex + 1) % 4
	return day6Directions[dirIndex], dirIndex
}

func solveDay6Part1(startPos []int, _map [][]bool) (int, map[int][]int) {
	visited := map[int][]int{}
	dirIndex := 0
	currentDir := day6Directions[dirIndex]
	currentPos := []int{startPos[0], startPos[1]}
	step := 0
	for true {
		if currentPos[0] < 0 || currentPos[0] >= len(_map) ||
			currentPos[1] < 0 || currentPos[1] >= len(_map[0]) {
			break
		}
		if _map[currentPos[0]][currentPos[1]] {
			currentPos[0] -= currentDir[0]
			currentPos[1] -= currentDir[1]
			currentDir, dirIndex = turnRight(dirIndex)
			continue
		}
		hash := utils.TwoDToOneD(currentPos[1], currentPos[0], len(_map[0]))
		if v, ok := visited[hash]; ok {
			if v[dirIndex] >= 0 {
				return -1, nil
			}
		} else {
			visited[hash] = []int{-1, -1, -1, -1}
		}
		visited[hash][dirIndex] = step
		currentPos[0] += currentDir[0]
		currentPos[1] += currentDir[1]
		step++
	}
	return len(visited), visited
}

func hasLoop(x, y, step, dirIndex int, originalVisited map[int][]int, _map [][]bool) bool {
	visited := map[int][]bool{}
	currentDir := day6Directions[dirIndex]
	currentPos := []int{y, x}
	for true {
		if currentPos[0] < 0 || currentPos[0] >= len(_map) ||
			currentPos[1] < 0 || currentPos[1] >= len(_map[0]) {
			break
		}
		if _map[currentPos[0]][currentPos[1]] ||
			(currentPos[0] == y && currentPos[1] == x) {
			currentPos[0] -= currentDir[0]
			currentPos[1] -= currentDir[1]
			currentDir, dirIndex = turnRight(dirIndex)
			continue
		}
		hash := utils.TwoDToOneD(currentPos[1], currentPos[0], len(_map[0]))
		if v, ok := visited[hash]; ok {
			if v[dirIndex] {
				return true
			}
		} else {
			if v, ok := originalVisited[hash]; ok {
				if v[dirIndex] >= 0 && v[dirIndex] < step {
					return true
				}
			}
			visited[hash] = []bool{false, false, false, false}
		}
		visited[hash][dirIndex] = true
		currentPos[0] += currentDir[0]
		currentPos[1] += currentDir[1]
	}
	return false
}

func solveDay6Part2(startPos []int, _map [][]bool, visited map[int][]int) int {
	keys := []int{}
	for k := range visited {
		keys = append(keys, k)
	}
	fn := func(i int) int {
		k := keys[i]
		x, y := utils.OneDTwoD(k, len(_map[0]))
		if x == startPos[1] && y == startPos[0] {
			return 0
		}
		dirs := visited[k]
		lowestDirIndex := -1
		for i := 0; i < 4; i++ {
			if dirs[i] >= 0 {
				if lowestDirIndex == -1 || dirs[i] < dirs[lowestDirIndex] {
					lowestDirIndex = i
				}
			}
		}

		if hasLoop(x, y, dirs[lowestDirIndex]-1, lowestDirIndex, visited, _map) {
			return 1
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(keys))
}

func Day6(input []string) []string {
	// TODO - perf: precalculate jumptable for all obstacles in the map
	// must support step preservation && always give valid next coords
	// need to be able to disable for jumps that include extra obstacles
	currentPos, _map := parseDay6Input(input)
	solution1, visited := solveDay6Part1(currentPos, _map)
	solution2 := solveDay6Part2(currentPos, _map, visited)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
