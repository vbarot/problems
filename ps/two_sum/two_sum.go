package two_sum

import (
	"fmt"
)

func Describe() {
	const desc = `##########
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
##########`
	fmt.Println(desc)
}

func Test(){
	test_1()
}

func test_1() {
	nums := []int{2, 7, 11, 15}
	target := 9
	out := twoSum(nums, target)
	fmt.Println(out)
}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i,v := range nums {
		compliment := target - v
		if val, exists := m[compliment]; exists {
			return []int{i,val}
		}
		m[v]=i
	}

	return nil
}
