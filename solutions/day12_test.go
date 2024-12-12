package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay12TestData(t *testing.T) {
	input := utils.GetTestInput(12)
	expected := []string{"1930"}
	actual := Day12(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
