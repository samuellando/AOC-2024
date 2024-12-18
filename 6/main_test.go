
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 4939 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 1434 {
        t.Fail()
    }
}
