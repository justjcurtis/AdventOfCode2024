package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay5TestInput(t *testing.T) {
	input := utils.GetTestInput(5)
	expected := []string{"143", "123"}
	actual := Day5(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestDay5(t *testing.T) {
	input := utils.GetInputForTest(5)
	expected := []string{"5955", "4030"}
	actual := Day5(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
