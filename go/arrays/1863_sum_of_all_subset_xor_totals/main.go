package main

func subsetXORSum(nums []int) int {
	total := 0
	XORSum(&nums, &total, 0, 0)
	return total
}

func XORSum(nums *[]int, total *int, i int, currentXOR int) {
	if i == len(*nums) {
		*total += currentXOR
		return
	} else {
		// without value
		XORSum(nums, total, i+1, currentXOR)

		// with value
		XORSum(nums, total, i+1, currentXOR^((*nums)[i]))
	}
}

func main() {
	nums := []int{1, 3}
	res := subsetXORSum(nums)
	println(res)
}
