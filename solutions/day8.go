package solutions

import (
	"AdventOfCode2024/utils"
	"strconv"
	"sync"
)

type AntennaMap struct {
	width  int
	height int
	coords [][][]int
}

func parseDay8(input []string) AntennaMap {
	antennaMap := AntennaMap{len(input[0]), len(input), [][][]int{}}
	runeMap := map[rune]int{}
	mu := sync.Mutex{}
	fn := func(i int) {
		line := input[i]
		for j, char := range line {
			if char == '.' {
				continue
			}
			mu.Lock()
			if _, ok := runeMap[char]; !ok {
				runeMap[char] = len(antennaMap.coords)
				antennaMap.coords = append(antennaMap.coords, [][]int{})
			}
			index := runeMap[char]
			antennaMap.coords[index] = append(antennaMap.coords[index], []int{j, i})
			mu.Unlock()
		}
	}
	utils.ParalleliseVoid(fn, len(input))
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
	maxHash := utils.TwoDToOneD(antennaMap.width-1, antennaMap.height-1, antennaMap.width)
	antinodeHashes := make([]bool, maxHash+1)
	antinodeCount := 0
	mu := sync.Mutex{}
	fn := func(index int) {
		coords := antennaMap.coords[index]
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				antinodes := getAntinodeCoords(coords[i], coords[j], antennaMap.width, antennaMap.height, getAll)
				for _, antinode := range antinodes {
					hash := utils.TwoDToOneD(antinode[1], antinode[0], antennaMap.width)
					mu.Lock()
					if !antinodeHashes[hash] {
						antinodeHashes[hash] = true
						antinodeCount++
					}
					mu.Unlock()
				}
			}
		}
	}
	utils.ParalleliseVoid(fn, len(antennaMap.coords))
	return antinodeCount
}

func Day8(input []string) []string {
	parsed := parseDay8(input)
	solution1 := solveDay8(parsed, false)
	solution2 := solveDay8(parsed, true)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
