package utils

import (
	"reflect"
	"testing"
)

func TestDivideSliceToSlices(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7}

	expectedResult := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7},
	}

	actualResult, _ := DivideSliceToSlices(slice, 2)

	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("failed equal slice")
	}

}

func TestRevertKeys(t *testing.T) {
	originalMap := map[int]int{2: 1, 3: 2}

	expectedResult := map[int]int{1: 2, 2: 3}

	actualResult, _ := RevertKeys(originalMap)

	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("failed revert map")
	}
}

func TestFilterSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	values := []int{4, 5}

	expectedResult := []int{1, 2, 3, 6, 7}

	actualResult := FilterSlice(slice, values)

	t.Log(actualResult)

	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("failed filter map")
	}

}
