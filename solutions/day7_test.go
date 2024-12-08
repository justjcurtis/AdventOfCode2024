package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay7TestInput(t *testing.T) {
	input := utils.GetTestInput(7)
	expected := []string{"3749", "11387"}
	actual := Day7(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
