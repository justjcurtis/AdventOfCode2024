package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay13TestInput(t *testing.T) {
	input := utils.GetTestInput(13)
	expected := []string{"480", "875318608908"}
	result := Day13(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Day13: expected %v, got %v", expected, result)
	}
}

func TestDay13TestInputPart2(t *testing.T) {
	input := utils.GetInputForTest(13)
	expected := []string{"34787", "85644161121698"}
	result := Day13(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Day13: expected %v, got %v", expected, result)
	}
}
