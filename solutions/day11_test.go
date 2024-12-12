package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay11TestInput(t *testing.T) {
	input := utils.GetTestInput(11)
	expected := []string{"55312"}
	result := Day11(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but was %v", expected, result)
	}
}

func TestDay11(t *testing.T) {
	input := utils.GetInputForTest(11)
	expected := []string{"183484"}
	result := Day11(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but was %v", expected, result)
	}
}
