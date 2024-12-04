package common

import (
	"bufio"
	"iter"
	"os"
	"strconv"
	"strings"
)

func AsInts(in iter.Seq[[]string]) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for line := range in {
			iline := make([]int, 0, len(line))
			for _, s := range line {
				iline = append(iline, Net(strconv.Atoi(s)))
			}
			if !yield(iline) {
				return
			}
		}
	}
}

func Input() []byte {
	return Net(os.ReadFile("input.txt"))
}

func InputLines() iter.Seq[[]string] {
	return func(yield func([]string) bool) {
		f := Net(os.Open("input.txt"))
		defer f.Close()
		s := bufio.NewScanner(f)
		for s.Scan() {
			ns := strings.Split(s.Text(), " ")
			line := make([]string, 0, 100)
			for _, ss := range ns {
				if ss == "" {
					continue
				}
				line = append(line, ss)
			}
			if !yield(line) {
				return
			}
		}
	}
}
