/*
Copyright © 2024 Jacson Curtis <justjcurtis@gmail.com>
*/
package main

import (
	"AdventOfCode2024/solutions"
	"AdventOfCode2024/utils"
	"flag"
	"fmt"
	"time"
)

type solution struct {
	day int
	fn  func([]string) []string
}

var SOLUTIONS = []solution{
	{1, solutions.Day1},
	{2, solutions.Day2},
}

func main() {
	runCount := flag.Int("n", 1, "Number of times to run each solution")
	minRun := flag.Bool("min", false, "Use the minimum run time instead of the average")
	singleDay := flag.Int("d", -1, "Run only the specified day")

	flag.Parse()

	if *singleDay > len(SOLUTIONS) {
		println("Invalid day specified")
		return
	}

	if *runCount < 1 {
		*runCount = 1
	}
	if *minRun && *runCount < 2 {
		*runCount = 5000
	}

	var totalTime time.Duration
	for d, solution := range SOLUTIONS {
		if *singleDay > -1 && *singleDay != d+1 {
			continue
		}
		minElapsed := time.Duration(0)
		input := utils.GetInput(solution.day)
		if *minRun {
			start := time.Now()
			for i := 0; i < *runCount-1; i++ {
				start = time.Now()
				solution.fn(input)
				elapsed := time.Since(start)
				if elapsed < minElapsed || minElapsed == 0 {
					minElapsed = elapsed
				}
			}
			results := solution.fn(input)
			totalTime += minElapsed
			utils.PrintResults(solution.day, results)
			fmt.Printf("Day %d took %s\n", solution.day, minElapsed)
		} else {
			start := time.Now()
			for i := 0; i < *runCount-1; i++ {
				solution.fn(input)
			}
			results := solution.fn(input)
			elapsed := time.Since(start)
			totalTime += elapsed
			utils.PrintResults(solution.day, results)
			fmt.Printf("Day %d took %s\n", solution.day, elapsed/time.Duration(*runCount))
		}
		if *singleDay == -1 {
			println()
		}
	}

	if *singleDay == -1 {
		println("=------ Total ------=")
		if *minRun {
			fmt.Printf("Total time: %s\n", totalTime)
		} else {
			fmt.Printf("Total time: %s\n", totalTime/time.Duration(*runCount))
		}
	}
	println("=-------------------=")
}
