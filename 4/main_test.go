
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 2483 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 1925 {
        t.Fail()
    }
}
