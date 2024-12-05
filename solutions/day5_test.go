package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay5TestInput(t *testing.T) {
	input := utils.GetTestInput(5)
	expected := []string{"143"}
	actual := Day5(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
