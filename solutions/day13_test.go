package solutions

import (
	"AdventOfCode2024/utils"
	"reflect"
	"testing"
)

func TestDay13TestInput(t *testing.T) {
	input := utils.GetTestInput(13)
	expected := []string{"480"}
	result := Day13(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Day13: expected %v, got %v", expected, result)
	}
}
