package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay1(t *testing.T) {
	input := utils.GetInputForTest(1)
	expected := []string{"2000468", "18567089"}
	actual := Day1(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
