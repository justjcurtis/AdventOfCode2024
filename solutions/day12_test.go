package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay12TestData(t *testing.T) {
	input := utils.GetTestInput(12)
	expected := []string{"1930", "1206"}
	actual := Day12(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestDay12(t *testing.T) {
	input := utils.GetInputForTest(12)
	expected := []string{"1465968", "897702"}
	actual := Day12(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
