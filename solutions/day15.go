package solutions

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strconv"
	"sync/atomic"
)

const EMPTY = 0
const WALL = 1
const BOX = 2
const ROBOT = 3
const UP = 0
const RIGHT = 1
const DOWN = 2
const LEFT = 3

func day15CharToInt(c byte) int {
	switch c {
	case '.':
		return EMPTY
	case '#':
		return WALL
	case 'O':
		return BOX
	case '@':
		return ROBOT
	case '^':
		return UP
	case '>':
		return RIGHT
	case 'v':
		return DOWN
	case '<':
		return LEFT
	}
	return -1
}

func parseDay15(input []string) ([]int, []int, int, int) {
	gridStr := []string{}
	breakIndex := 0
	for i, line := range input {
		if len(line) == 0 {
			breakIndex = i
			break
		}
		gridStr = append(gridStr, line)
	}
	parsed := make([]int, len(gridStr)*len(gridStr[0]))
	fn := func(i int) {
		x, y := utils.OneDTwoD(i, len(input[0]))
		parsed[i] = day15CharToInt(input[y][x])
	}
	utils.ParalleliseVoid(fn, len(parsed))
	instructions := []int{}
	for i := breakIndex + 1; i < len(input); i++ {
		for _, c := range input[i] {
			instructions = append(instructions, day15CharToInt(byte(c)))
		}
	}
	return parsed, instructions, len(gridStr[0]), len(gridStr)
}

func copyGrid(grid []int) []int {
	newGrid := make([]int, len(grid))
	copy(newGrid, grid)
	return newGrid
}

func handlePush(grid []int, x, y, w, h, dir int, doubleWide bool) bool {
	hash := utils.TwoDToOneD(x, y, w)
	if grid[hash] == EMPTY {
		return true
	}
	if grid[hash] == WALL {
		return false
	}
	if grid[hash] == BOX || grid[hash] == -BOX {
		nextX, nextY := x, y
		switch dir {
		case UP:
			nextY--
		case RIGHT:
			nextX++
		case DOWN:
			nextY++
		case LEFT:
			nextX--
		}
		if nextX < 0 || nextX >= w || nextY < 0 || nextY >= h {
			return false
		}
		if doubleWide && (dir == UP || dir == DOWN) {
			otherNextX := nextX
			if grid[hash] == BOX {
				otherNextX--
			} else {
				otherNextX++
			}
			if otherNextX < 0 || otherNextX >= w {
				return false
			}
			otherNextHash := utils.TwoDToOneD(otherNextX, nextY, w)
			if !handlePush(grid, otherNextX, nextY, w, h, dir, doubleWide) {
				return false
			}
			otherCurrentHash := utils.TwoDToOneD(otherNextX, y, w)
			grid[otherNextHash] = grid[otherCurrentHash]
			grid[otherCurrentHash] = EMPTY
		}
		if !handlePush(grid, nextX, nextY, w, h, dir, doubleWide) {
			return false
		}
		nextHash := utils.TwoDToOneD(nextX, nextY, w)
		grid[nextHash] = grid[hash]
		grid[hash] = EMPTY
		return true
	}
	return true
}

func handleMove(grid []int, r, w, h, dir int, doubleWide bool) int {
	x, y := utils.OneDTwoD(r, w)
	switch dir {
	case UP:
		y--
	case RIGHT:
		x++
	case DOWN:
		y++
	case LEFT:
		x--
	default:
		return r
	}
	if x < 0 || x >= w || y < 0 || y >= h {
		return r
	}
	var originalGrid []int
	if doubleWide {
		originalGrid = make([]int, len(grid))
		copy(originalGrid, grid)
	}
	if handlePush(grid, x, y, w, h, dir, doubleWide) {
		nextR := utils.TwoDToOneD(x, y, w)
		grid[nextR] = ROBOT
		grid[r] = EMPTY
		return nextR
	} else if doubleWide {
		copy(grid, originalGrid)
	}
	return r
}

func printDay15Grid(grid []int, w, h int, doubleWide bool) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			hash := utils.TwoDToOneD(j, i, w)
			switch grid[hash] {
			case EMPTY:
				fmt.Print(".")
			case WALL:
				fmt.Print("#")
			case BOX:
				if doubleWide {
					fmt.Print("]")
				} else {
					fmt.Print("O")
				}
			case -BOX:
				if doubleWide {
					fmt.Print("[")
				}
			case ROBOT:
				fmt.Print("@")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func solveDay15Part1(parsed, instructions []int, w, h int) int {
	grid := copyGrid(parsed)
	r := -1
	for i, v := range grid {
		if v == ROBOT {
			r = i
			break
		}
	}
	for _, instruction := range instructions {
		r = handleMove(grid, r, w, h, instruction, false)
	}

	result := int64(0)
	fn := func(i int) {
		if grid[i] == BOX {
			x, y := utils.OneDTwoD(i, w)
			atomic.AddInt64(&result, int64((y*100)+x))
		}
	}
	utils.ParalleliseVoid(fn, w*h)
	return int(result)
}

func resizeGrid(grid []int, w, h int) []int {
	newGrid := make([]int, w*2*h)
	index := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			hash := utils.TwoDToOneD(j, i, w)
			switch grid[hash] {
			case EMPTY:
				newGrid[index] = EMPTY
				newGrid[index+1] = EMPTY
			case WALL:
				newGrid[index] = WALL
				newGrid[index+1] = WALL
			case BOX:
				newGrid[index] = -BOX
				newGrid[index+1] = BOX
			case ROBOT:
				newGrid[index] = ROBOT
				newGrid[index+1] = EMPTY
			}
			index += 2
		}
	}
	return newGrid
}

func solveDay15Part2(parsed, instructions []int, w, h int) int {
	grid := resizeGrid(copyGrid(parsed), w, h)
	r := -1
	for i, v := range grid {
		if v == ROBOT {
			r = i
			break
		}
	}
	for _, instruction := range instructions {
		r = handleMove(grid, r, w*2, h, instruction, true)
	}
	result := int64(0)
	fn := func(i int) {
		if grid[i] != -BOX {
			return
		}
		x, y := utils.OneDTwoD(i, w*2)
		atomic.AddInt64(&result, int64((y*100)+x))
	}
	utils.ParalleliseVoid(fn, w*h*2)
	return int(result)
}

func Day15(input []string) []string {
	parsed, instructions, w, h := parseDay15(input)
	solution1 := solveDay15Part1(parsed, instructions, w, h)
	solution2 := solveDay15Part2(parsed, instructions, w, h)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
