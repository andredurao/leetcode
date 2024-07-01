//

package main

func main() {
	arr := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	println(threeConsecutiveOdds(arr))
}

func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	count, i, lastEven := 0, 0, -1
	for i <= len(arr)-3 {
		count = 0
		if arr[i]%2 == 1 {
			count++
		} else {
			lastEven = i
		}
		if arr[i+1]%2 == 1 {
			count++
		} else {
			lastEven = i + 1
		}
		if arr[i+2]%2 == 1 {
			count++
		} else {
			lastEven = i + 2
		}

		if count == 3 {
			return true
		}
		i = lastEven + 1
	}
	return false
}

// func threeConsecutiveOdds(arr []int) bool {
// 	if len(arr) < 3 {
// 		return false
// 	}
// 	count, i, lastEven := 0, 0, -1
// 	for i <= len(arr)-3 {
// 		count = 0
// 		checkOdd(&arr, &count, &lastEven, i)
// 		checkOdd(&arr, &count, &lastEven, i+1)
// 		checkOdd(&arr, &count, &lastEven, i+2)
// 		i = lastEven + 1
// 	}
// 	return false
// }

// func checkOdd(arr *[]int, count *int, lastEven *int, i int) {
// 	if (*arr)[i]%2 == 1 {
// 		(*count)++
// 	} else {
// 		*lastEven = i
// 	}
// }
