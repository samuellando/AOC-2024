
package main
import (
    "testing"
)

func TestPart1(t *testing.T) {
    if Part1() != 336 {
        t.Fail()
    }
}

func TestPart2(t *testing.T) {
    if Part2() != 758890600222015 {
        t.Fail()
    }
}
