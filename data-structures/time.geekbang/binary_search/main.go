package main

func BinarySearch(a []int, value int) int {
	low, high := 0, len(a)-1

	for low <= high {
		mid := low + ((high - low) >> 1)

		if a[mid] == value {
			return mid
		}
		if value > a[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchRecusion(a []int, value int) int {
	var bSearch func([]int, int, int, int) int

	bSearch = func(a []int, low int, high int, value int) int {
		if low > high {
			return -1
		}

		mid := low + ((high - low) >> 1)
		if value == a[mid] {
			return mid
		}
		if value > a[mid] {
			return bSearch(a, mid+1, high, value)
		} else {
			return bSearch(a, low, mid-1, value)
		}
	}

	return bSearch(a, 0, len(a), value)
}