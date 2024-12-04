package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay4TestInput(t *testing.T) {
	input := utils.GetTestInput(4)
	expected := []string{"18", "9"}
	result := Day4(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestDay4(t *testing.T) {
	input := utils.GetInputForTest(4)
	expected := []string{"2524", "1873"}
	result := Day4(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
