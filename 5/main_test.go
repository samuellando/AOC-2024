
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 5391 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 6142 {
        t.Fail()
    }
}
