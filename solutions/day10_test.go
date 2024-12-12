package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay10TestInput(t *testing.T) {
	input := utils.GetTestInput(10)
	expected := []string{"36", "81"}
	result := Day10(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestDay10(t *testing.T) {
	input := utils.GetInputForTest(10)
	expected := []string{"468", "966"}
	result := Day10(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}
