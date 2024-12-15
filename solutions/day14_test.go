package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay14TestInput(t *testing.T) {
	input := utils.GetTestInput(14)
	expected := []string{"12", "0"}
	result := Day14(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Day14: %v, expected %v", result, expected)
	}
}

func TestDay14(t *testing.T) {
	input := utils.GetInputForTest(14)
	expected := []string{"218433348", "6512"}
	result := Day14(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Day14: %v, expected %v", result, expected)
	}
}
