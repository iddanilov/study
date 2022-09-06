package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 3))

}

func binarySearch(nums []int, findValue int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == findValue {
			return true
		}
		if nums[m] < findValue {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
