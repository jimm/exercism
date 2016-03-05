package binarysearch

import "fmt"

func SearchInts(ints []int, key int) int {
	idx := mainSearch(ints, key)
	if idx < len(ints) {
		for idx > 0 {
			if ints[idx-1] == ints[idx] {
				idx--
			} else {
				break
			}
		}
	}
	return idx
}

func mainSearch(ints []int, key int) int {
	min := 0
	max := len(ints) - 1
	if min >= max {
		return 0
	}
	mid := 0
	for {
		if max-min == 1 {
			if key <= ints[min] {
				return min
			} else if key > ints[max] {
				return max + 1
			}
			return max
		}
		mid = min + (max-min)/2
		if ints[mid] == key {
			return mid
		} else if ints[mid] < key {
			min = mid
		} else {
			max = mid
		}
	}
}

func Message(ints []int, key int) string {
	if len(ints) == 0 {
		return "slice has no values"
	}

	idx := SearchInts(ints, key)
	if idx >= len(ints) {
		return fmt.Sprintf("%d > all %d values", key, len(ints))
	}

	found := ints[idx] == key
	if found {
		if idx == 0 {
			return fmt.Sprintf("%d found at beginning of slice", key)
		} else if idx == len(ints)-1 {
			return fmt.Sprintf("%d found at end of slice", key)
		}
		return fmt.Sprintf("%d found at index %d", key, idx)
	}

	if idx == 0 && key < ints[0] {
		return fmt.Sprintf("%d < all values", key)
	}
	return fmt.Sprintf("%d > %d at index %d, < %v at index %v", key,
		ints[idx-1], idx-1,
		ints[idx], idx)
}
