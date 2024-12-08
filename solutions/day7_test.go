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

func TestDay7(t *testing.T) {
	input := utils.GetInputForTest(7)
	expected := []string{"2299996598890", "362646859298554"}
	actual := Day7(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
