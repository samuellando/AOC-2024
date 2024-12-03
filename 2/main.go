package main

import (
	"advent/common"
	"fmt"
	"iter"
	"slices"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func Part1() int {
	count := 0
	for line := range common.AsInts(common.InputLines()) {
		if isSafe(line) {
			count++
		}
	}
    return count
}
func Part2() int {
	count := 0
	for line := range common.AsInts(common.InputLines()) {
		if common.Any(common.Map(generateWithRemoval(line, 1), isSafe)) {
			count++
		}
	}
    return count
}

func generateWithRemoval(report []int, remove int) iter.Seq[[]int] {
	if remove <= 0 {
		return func(yield func([]int) bool) {
			yield(report)
			return
		}
	}
	return func(yield func([]int) bool) {
		for i := range len(report) {
			ir := slices.Concat(report[:i], report[i+1:])
			for sub := range generateWithRemoval(ir, remove-1) {
				if !yield(sub) {
					return
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

func pairs(line []int) iter.Seq[[]int] {
	return func(yeild func([]int) bool) {
		for i := 1; i < len(line); i++ {
			if !yeild([]int{line[i-1], line[i]}) {
				return
			}
		}
	}
}
