
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 2031679 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 19678534 {
        t.Fail()
    }
}
