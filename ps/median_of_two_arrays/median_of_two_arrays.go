package median_of_two_arrays

import (
	"fmt"
)

func Describe() {
	const desc = `##########
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:
	nums1 = [1, 3]
	nums2 = [2]

	The median is 2.0

Example 2:
	nums1 = [1, 2]
	nums2 = [3, 4]

	The median is (2 + 3)/2 = 2.5

##########`
	fmt.Println(desc)
}

func Test(){
	test_1()
}

func test_1() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	out := findMedianSortedArrays(nums1, nums2)
	fmt.Println(out)
}

func test_2() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	out := findMedianSortedArrays(nums1, nums2)
	fmt.Println(out)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}