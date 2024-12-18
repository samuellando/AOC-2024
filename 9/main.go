package main

import (
	"advent/common"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Part 1:")
	fmt.Println(Part1())
	fmt.Println("Part 2:")
	fmt.Println(Part2())
}

func Part1() int {
	input := string(common.Input())
	disk := make([]int, 0)
	// Get a current reprisentation of the disk
	for i, c := range input {
		if c == '\n' {
			break
		}
		v := common.Net(strconv.Atoi(string(c)))
		for range v {
			if i%2 == 0 {
				disk = append(disk, i/2)
			} else {
				disk = append(disk, -1)
			}
		}
	}
	// Defragment
	for ip, rp := 0, len(disk)-1; ip < rp; {
		if disk[rp] == -1 {
			rp--
		} else if disk[ip] == -1 {
			disk[ip] = disk[rp]
			disk[rp] = -1
			ip++
			rp--
		} else {
			ip++
		}
	}
	// Calculate
	res := 0
	for i, c := range disk {
		if c == -1 {
			break
		}
		res += i * c
	}
	return res
}

type file struct {
	id     int
	blocks int
}

func Part2() int {
	input := string(common.Input())
	disk := make([]*file, 0)
	// Get a current reprisentation of the disk
	for i, c := range input {
		if c == '\n' {
			break
		}
		v := common.Net(strconv.Atoi(string(c)))
		if i%2 == 0 {
			f := file{i / 2, v}
			disk = append(disk, &f)
		} else {
			f := file{-1, v}
			disk = append(disk, &f)
		}
	}
	// Defragment
	for rp := len(disk)-1; rp >= 0; {
		if disk[rp].id == -1 {
			rp--
		} else {
			for p := 0; p < rp; p++ {
                if disk[p].id != -1 {
                    continue
                }
				remainder := disk[p].blocks - disk[rp].blocks
				if remainder >= 0 {
                    disk[p] = disk[rp]
					disk[rp] = &file{-1, disk[rp].blocks}
					disk = slices.Insert(disk, p+1, &file{-1, remainder})
                    break
				}
			}
					rp--
		} 
	}
	// Calculate
	res := 0
	di := 0
	for _, f := range disk {
		if f.id == -1 {
            di += f.blocks
            continue
		}
		for range f.blocks {
			res += di * f.id
			di++
		}
	}
	return res
}
