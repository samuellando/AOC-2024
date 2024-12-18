
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 6330095022244 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 6359491814941 {
        t.Fail()
    }
}
