
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 479 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 531 {
        t.Fail()
    }
}
