package common

import (
    "os"
    "bufio"
    "strings"
    "strconv"
)

func GetInputs() [][]int {
	f, err := os.Open("input.txt")
    defer f.Close()
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
    lines := make([][]int, 0, 1000)

	for s.Scan() {
		ns := strings.Split(s.Text(), " ")
		line := make([]int, 0, 100)
		for _, ss := range ns {
            if ss == "" {
                continue
            }
			value, err := strconv.Atoi(ss)
			if err != nil {
				panic(err)
			}
            line = append(line, value)
		}
        lines = append(lines, line)
	}
    return lines
}
