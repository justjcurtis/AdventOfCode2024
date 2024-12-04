package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay3(t *testing.T) {
	input := utils.GetInputForTest(3)
	expected := []string{"178794710", "76729637"}
	actual := Day3(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
