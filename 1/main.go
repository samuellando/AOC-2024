package main

import (
	"fmt"
	"sort"
    "advent/common"
)

func main() {
    fmt.Println("Part 1:")
    fmt.Println(Part1())
    fmt.Println("Part 2:")
    fmt.Println(Part2())
}

func Part1() int {
    r := make([]int, 0)
    l := make([]int, 0)
    for line := range(common.AsInts(common.InputLines())) {
        r = append(r, line[0])
        l = append(l, line[len(line) - 1])
    }
    sort.Sort(sort.IntSlice(r))
    sort.Sort(sort.IntSlice(l))
    sum := 0
    for i := 0; i < len(r); i++ {
        sum += common.Abs(r[i] - l[i])
    }
    return sum
}

func Part2() int {
    r := make(map[int]int)
    l := make([]int, 0)
    for line := range(common.AsInts(common.InputLines())) {
        ri := line[0]
        li := line[len(line) - 1]
        if _, ok := r[ri]; ok {
            r[ri] += 1
        } else {
            r[ri] = 1
        }
        l = append(l, li)
    }
    sum := 0
    for i := 0; i < len(l); i++ {
        if val, ok := r[l[i]]; ok {
            sum += val * l[i]
        } 
    }
    return sum
}
