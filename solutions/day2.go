package solutions

import (
	"AdventOfCode2024/utils"
	"strconv"
	"strings"
)

func parseDay2Line(line string) []int {
	levs := strings.Split(line, " ")
	nums := []int{}
	for _, lev := range levs {
		num, _ := strconv.Atoi(lev)
		nums = append(nums, num)
	}
	return nums
}

func parseDay2(input []string) [][]int {
	reports := make([][]int, len(input))
	fn := func(i int) {
		report := parseDay2Line(input[i])
		reports[i] = report
	}
	utils.ParalleliseVoid(fn, len(input))
	return reports
}

func isReportSafe(report []int) (bool, int) {
	increasing := true
	for i, b := range report {
		if i == 0 {
			c := report[1]
			if c > b {
				increasing = true
			}
			if c < b {
				increasing = false
			}
			continue
		}
		a := report[i-1]
		if a == b {
			return false, i
		}
		if increasing && a > b {
			return false, i
		}
		if !increasing && a < b {
			return false, i
		}
		diff := a - b
		if diff < 0 {
			diff = -diff
		}
		if diff > 3 {
			return false, i
		}

	}
	return true, -1
}

func solveDay2Part1(reports [][]int) int {
	fn := func(i int) int {
		isSafe, _ := isReportSafe(reports[i])
		if isSafe {
			return 1
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(reports))
}

func solveDay2Part2(reports [][]int) int {
	fn := func(i int) int {
		isSafe, badIndex := isReportSafe(reports[i])
		if isSafe {
			return 1
		}
		for skip := badIndex - 2; skip <= badIndex+1; skip++ {
			if skip < 0 || skip >= len(reports[i]) {
				continue
			}
			skippedReport := append([]int{}, reports[i][:skip]...)
			skippedReport = append(skippedReport, reports[i][skip+1:]...)
			isSafe, _ := isReportSafe(skippedReport)
			if isSafe {
				return 1
			}
		}
		return 0
	}
	return utils.Parallelise(utils.IntAcc, fn, len(reports))
}

func Day2(input []string) []string {
	parsedReports := parseDay2(input)
	solution1 := solveDay2Part1(parsedReports)
	solution2 := solveDay2Part2(parsedReports)
	return []string{strconv.Itoa(solution1), strconv.Itoa(solution2)}
}
