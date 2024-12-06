package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay6TestInput(t *testing.T) {
	input := utils.GetTestInput(6)
	expected := []string{"41", "6"}
	actual := Day6(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
