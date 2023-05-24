// https://leetcode.com/problems/find-in-mountain-array/description/

package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{1, 5, 2}
	target := 0
	valuesMap := make(map[int]int)
	peak := findPeakIndex(array, valuesMap)
	fmt.Println(peak, array[peak])
	searchLeft := lookupLeft(array, target, 0, peak+1, valuesMap)
	fmt.Println("searchLeft", searchLeft)
	searchRight := lookupRight(array, target, peak+1, len(array))
	fmt.Println("searchRight", searchRight)
}

func get(array []int, valuesMap map[int]int, index int) int {
	value, ok := valuesMap[index]
	if ok {
		return value
	}
	value = array[index]
	valuesMap[index] = value
	return value
}

func findPeakIndex(array []int, valuesMap map[int]int) int {
	l := 1
	r := len(array) - 1

	fmt.Printf("%v\n", array)
	for {
		m := (l + r) / 2
		fmt.Println("l", l, "r", r, "m", m)
		cur := get(array, valuesMap, m)
		prev := get(array, valuesMap, m-1)
		next := get(array, valuesMap, m+1)
		if (prev < cur) && (next < cur) {
			return m
		} else {
			if cur < next {
				l = m + 1
			} else {
				r = m
			}
		}
	}
}

func lookupLeft(array []int, target int, l int, r int, valuesMap map[int]int) int {
	for l < r {
		m := int(math.Floor(float64(l+r) / 2.0))
		fmt.Println("l", l, "r", r, "m", m, "array[m]", get(array, valuesMap, m))
		if get(array, valuesMap, m) < target {
			l = m + 1
		} else {
			r = m
		}
	}

	if get(array, valuesMap, l) == target {
		return l
	}
	return -1
}

func lookupRight(array []int, target int, l int, r int) int {
	// binary search leftmost
	for l < r {
		m := int(math.Floor(float64(l+r) / 2.0))
		fmt.Println("l", l, "r", r, "m", m, "array[m]", array[m])
		if array[m] > target {
			l = m + 1
		} else {
			r = m
		}
	}
	fmt.Println("out l", l)

	if l < len(array) && array[l] == target {
		return l
	}
	return -1
}

func searchLeftMost(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		m := int(math.Floor(float64(l+r) / 2.0))
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// Solution posted on leetcode
// TODO: create struct
// TODO: fix code to run locally
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	len := mountainArr.length()
	valuesMap := make(map[int]int)
	peak := findPeakIndex(mountainArr, len, valuesMap)
	peakValue := get(mountainArr, valuesMap, peak)

	// fmt.Println("peak", peak, get(mountainArr, valuesMap, peak))
	if peakValue == target {
		return peak
	}

	left := lookupLeft(mountainArr, target, 0, peak+1, valuesMap)
	if left != -1 {
		leftValue := get(mountainArr, valuesMap, left)
		if leftValue == target {
			// fmt.Println("found")
			return left
		}
	}

	right := lookupRight(mountainArr, target, peak+1, len, valuesMap)
	if right != -1 {
		rightValue := get(mountainArr, valuesMap, right)
		if rightValue == target {
			// fmt.Println("found")
			return right
		}
	}

	return -1
}

// cache get requests
func get(mountainArr *MountainArray, valuesMap map[int]int, index int) int {
	value, ok := valuesMap[index]
	if ok {
		return value
	}
	value = mountainArr.get(index)
	valuesMap[index] = value
	return value
}

func findPeakIndex(mountainArr *MountainArray, len int, valuesMap map[int]int) int {
	l := 1
	r := len - 1

	for {
		m := (l + r) / 2
		cur := get(mountainArr, valuesMap, m)
		prev := get(mountainArr, valuesMap, m-1)
		next := get(mountainArr, valuesMap, m+1)
		if (prev < cur) && (next < cur) {
			return m
		} else {
			if cur < next {
				l = m + 1
			} else {
				r = m
			}
		}
	}
}

func lookupLeft(mountainArr *MountainArray, target int, l int, r int, valuesMap map[int]int) int {
	for l < r {
		m := int(math.Floor(float64(l+r) / 2.0))
		if get(mountainArr, valuesMap, m) < target {
			l = m + 1
		} else {
			r = m
		}
	}

	if get(mountainArr, valuesMap, l) == target {
		return l
	}
	return -1
}

func lookupRight(mountainArr *MountainArray, target int, l int, r int, valuesMap map[int]int) int {
	len := r
	for l < r {
		m := int(math.Floor(float64(l+r) / 2.0))
		if get(mountainArr, valuesMap, m) > target {
			l = m + 1
		} else {
			r = m
		}
	}

	if l < len && get(mountainArr, valuesMap, l) == target {
		return l
	}
	return -1
}
