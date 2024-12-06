package solutions

import (
	"reflect"
	"testing"
)

func TestDay6TestInput(t *testing.T) {
	input := []string{}
	expected := []string{"41"}
	actual := Day6(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
