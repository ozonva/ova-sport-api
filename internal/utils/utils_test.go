package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestDivideSliceToSlices(t *testing.T) {
	testSet := []struct {
		inputSlice     []int
		inputBatchSize int
		expected       [][]int
		expectedErr    error
	}{
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}, nil},
		{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}, nil},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2, 3}}, nil},
		{[]int{1, 2, 3}, 0, nil, errors.New("chunk size can't be less 1")},
		{[]int{1, 2}, 3, [][]int{{1, 2}}, nil},
		{[]int{1, 2, 3}, 2, [][]int{{1, 2}, {3}}, nil},
		{[]int{}, 2, nil, errors.New("slice length can't be 0")},
		{nil, 1, nil, nil},
	}

	var actualResult [][]int
	var actualError error
	for i, test :=  range testSet {
		actualResult, actualError = DivideSliceToSlices(test.inputSlice, test.inputBatchSize)

		if !reflect.DeepEqual(actualResult, test.expected) {
			t.Errorf("Failed slice test %d. Wrong result", i)
		}

		if !reflect.DeepEqual(actualError, test.expectedErr) {
			t.Errorf("Failed slice test %d. Wrong error", i)
		}
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
