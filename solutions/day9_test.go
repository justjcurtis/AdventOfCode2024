package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay9TestInput(t *testing.T) {
	input := utils.GetTestInput(9)
	expected := []string{"1928", "2858"}
	actual := Day9(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}

func TestDay9(t *testing.T) {
	input := utils.GetInputForTest(9)
	expected := []string{"6301895872542", "6323761685944"}
	actual := Day9(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
