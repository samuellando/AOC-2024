package main

import (
	"advent/common"
	"fmt"
	"slices"
)

func main() {
    fmt.Println("Part 1:")
	part1()
    fmt.Println("Part 2:")
	part2()
}

func part1() {
	lines := common.GetInputs()
	count := 0
	for _, line := range lines {
		if isSafe(line) {
			count++
		}
	}
	fmt.Println(count)
}
func part2() {
	lines := common.GetInputs()
	count := 0
	for _, line := range lines {
		if common.Any(common.Map(generateWithRemoval(line, 1), isSafe)) {
			count++
		}
	}
	fmt.Println(count)
}

func generateWithRemoval(report []int, remove int) func() ([]int, bool) {
	if remove < 0 {
		return func() ([]int, bool) {
			return nil, false
		}
	}
	i := 0
	inner := generateWithRemoval(report, -1)
	return func() ([]int, bool) {
		if remove == 0 {
			if i == 0 {
				i++
				return report, true
			} else {
				return nil, false
			}
		} else {
			iv, ok := inner()
			if ok {
				return iv, true
			} else {
				if i < len(report) {
					ir := slices.Concat(report[:i], report[i+1:])
					inner = generateWithRemoval(ir, remove-1)
					i++
					v, ok := inner()
					if !ok {
						panic("huh")
					}
					return v, ok
				} else {
					return nil, false
				}
			}
		}
	}
}

func isSafe(report []int) bool {
	isInc := func(v []int) bool { return v[0] < v[1] }
	isDec := func(v []int) bool { return v[0] > v[1] }
	diffBelow := func(v []int) bool { return common.Abs(v[0]-v[1]) < 4 }
	return (common.All(common.Map(pairs(report), isInc)) ||
		common.All(common.Map(pairs(report), isDec))) &&
		common.All(common.Map(pairs(report), diffBelow))
}

func pairs(line []int) func() ([]int, bool) {
	i := 0
	return func() ([]int, bool) {
		if i < len(line)-1 {
			i++
			return []int{line[i-1], line[i]}, true
		} else {
			return nil, false
		}
	}
}
