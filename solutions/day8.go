package solutions

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strconv"
	"sync"
)

type AntennaMap struct {
	width  int
	height int
	coords map[rune][][]int
}

func printMapWithAntinodes(antennaMap AntennaMap, antinodeHashes *sync.Map) {
	_map := make([][]rune, antennaMap.height)
	for i := 0; i < antennaMap.height; i++ {
		_map[i] = make([]rune, antennaMap.width)
	}
	for i := 0; i < antennaMap.height; i++ {
		for j := 0; j < antennaMap.width; j++ {
			hash := utils.TwoDToOneD(j, i, antennaMap.width)
			if _, ok := antinodeHashes.Load(hash); ok {
				_map[i][j] = 'X'
			} else {
				_map[i][j] = '.'
			}
		}
	}
	for k, coords := range antennaMap.coords {
		for _, coord := range coords {
			_map[coord[1]][coord[0]] = k
		}
	}
	for i := 0; i < antennaMap.height; i++ {
		fmt.Println(string(_map[i]))
	}
}

func parseDay8(input []string) AntennaMap {
	antennaMap := AntennaMap{len(input[0]), len(input), make(map[rune][][]int)}
	for i, line := range input {
		for j, char := range line {
			if char == '.' {
				continue
			}
			if antennaMap.coords[char] == nil {
				antennaMap.coords[char] = [][]int{}
			}
			antennaMap.coords[char] = append(antennaMap.coords[char], []int{j, i})
		}
	}
	return antennaMap
}

func getAntinodeCoords(a []int, b []int, w, h int, getAll bool) [][]int {
	coords := [][]int{}
	for n := 2; n < 3 || getAll; n++ {
		newCoord := []int{n*a[0] - (n-1)*b[0], n*a[1] - (n-1)*b[1]}
		if newCoord[0] < 0 || newCoord[0] >= h || newCoord[1] < 0 || newCoord[1] >= w {
			break
		}
		coords = append(coords, newCoord)
	}
	for n := 2; n < 3 || getAll; n++ {
		newCoord := []int{n*b[0] - (n-1)*a[0], n*b[1] - (n-1)*a[1]}
		if newCoord[0] < 0 || newCoord[0] >= h || newCoord[1] < 0 || newCoord[1] >= w {
			break
		}
		coords = append(coords, newCoord)
	}

	if getAll {
		coords = append(coords, a)
		coords = append(coords, b)
	}
	return coords
}

func solveDay8(antennaMap AntennaMap, getAll bool) int {
	antinodeHashes := sync.Map{}
	keys := []rune{}
	for key := range antennaMap.coords {
		keys = append(keys, key)
	}
	fn := func(index int) {
		coords := antennaMap.coords[keys[index]]
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				antinodes := getAntinodeCoords(coords[i], coords[j], antennaMap.width, antennaMap.height, getAll)
				for _, antinode := range antinodes {
					hash := utils.TwoDToOneD(antinode[1], antinode[0], antennaMap.width)
					antinodeHashes.Store(hash, true)
				}
			}
		}
	}
	utils.ParalleliseVoid(fn, len(keys))
	return utils.SynMapLength(&antinodeHashes)
}

func Day8(input []string) []string {
	parsed := parseDay8(input)
	solution1 := solveDay8(parsed, false)
	solution2 := solveDay8(parsed, true)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
