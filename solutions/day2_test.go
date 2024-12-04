package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay2(t *testing.T) {
	input := utils.GetInputForTest(2)
	expected := []string{"483", "528"}
	actual := Day2(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
