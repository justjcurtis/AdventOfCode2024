package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay15TestInput(t *testing.T) {
	input := utils.GetTestInput(15)
	expected := []string{"10092", "9021"}
	result := Day15(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestDay15TestInput was incorrect, got: %v, want: %v.", result, expected)
	}
}
