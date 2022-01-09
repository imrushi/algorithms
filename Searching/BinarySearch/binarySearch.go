package binarysearch

// return the index
// return -1 if it does not exist
func BinarySearch(arr []int, target int) int {
	var start int = 0
	var end int = len(arr) - 1

	for start <= end {
		// find the middle element
		// mid  := (start + end) / 2 // If we use this formula if start and end value or sum of both will execed the limit of int range
		mid := start + (end-start)/2

		if target < arr[mid] {
			//check on left
			end = mid - 1
		} else if target > arr[mid] {
			// check on right
			start = mid + 1
		} else {
			// ans found
			return mid
		}
	}

	return -1
}

// return -1 if it does not exist
func OrderAgnosticBS(arr []int, target int) int {
	var start int = 0
	var end int = len(arr) - 1

	// find whether the array is sorted in ascending or descending
	var isAsc bool = arr[start] < arr[end]

	for start <= end {
		// find the middle element
		// mid  := (start + end) / 2 // If we use this formula if start and end value or sum of both will execed the limit of int range
		mid := start + (end-start)/2

		if arr[mid] == target {
			return mid
		}

		if isAsc {

			if target < arr[mid] {
				//check on left
				end = mid - 1
			} else if target > arr[mid] {
				// check on right
				start = mid + 1
			}
		} else {
			if target > arr[mid] {
				//check on right
				end = mid - 1
			} else if target < arr[mid] {
				// check on left
				start = mid + 1
			}
		}
	}
	return -1
}
