package challenges

import "errors"

func mergeSortedArrays(arr1 []int, arr2 []int) ([]int, error) {
	if arr1 == nil || arr2 == nil {
		return nil, errors.New("no empty arrays")
	}

	var mergedArray []int
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		numberFromFirstArray := arr1[i]
		numberFromSecondArray := arr2[j]
		if numberFromFirstArray < numberFromSecondArray {
			mergedArray = append(mergedArray, numberFromFirstArray)
			i++
		} else {
			mergedArray = append(mergedArray, numberFromSecondArray)
			j++
		}
	}

	for ; i < len(arr1); i++ {
		numberFromFirstArray := arr1[i]
		mergedArray = append(mergedArray, numberFromFirstArray)
	}

	for ; j < len(arr2); j++ {
		numberFromSecondArray := arr2[j]
		mergedArray = append(mergedArray, numberFromSecondArray)
	}

	return mergedArray, nil
}
