package utils

import "errors"

func DivideSliceToSlices(slice []int, chunkSize int) ([][]int, error){
	if chunkSize < 1 {
		return nil, errors.New("chunk size can't be less 1")
	}

	if len(slice) <= 0 {
		return nil, errors.New("slice length can't ")
	}
	result := make([][]int, 0, len(slice) / chunkSize + 1)

	for chunkSize < len(slice) {
		result = append(result, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}
	result = append(result, slice)

	return result, nil
}

func RevertKeys(sourceMap map[int]int) (map[int]int, error) {
	result := make(map[int]int)

	for key, value := range sourceMap {
		if _, ok := result[value]; ok {
			return nil, errors.New("key is not unique")
		}
		result[value] = key;
	}

	return result, nil
}

func FilterSlice(slice []int, values []int) []int {
	result := make([]int, 0)

	set := make(map[int]bool)
	for index := range values {
		set[values[index]] = true
	}

	for _, value := range slice {
		if !set[value] {
			result = append(result, value)
		}
	}

	return result
}