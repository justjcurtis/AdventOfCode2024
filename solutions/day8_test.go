package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay8TestInput(t *testing.T) {
	input := utils.GetTestInput(8)
	expected := []string{"14", "34"}
	actual := Day8(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestDay8(t *testing.T) {
	input := utils.GetInputForTest(8)
	expected := []string{"247", "861"}
	actual := Day8(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
