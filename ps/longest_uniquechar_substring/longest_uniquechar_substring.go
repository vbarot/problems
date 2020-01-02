package longest_uniquechar_substring

import (
	"fmt"
)

func Describe() {
	const desc = `##########
Given a string, find the length of the longest substring without repeating characters.

Example 1:
	Input: "abcabcbb"
	Output: 3 
	Explanation: The answer is "abc", with the length of 3. 

Example 2:
	Input: "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.
Example 3:
	Input: "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
##########`
	fmt.Println(desc)
}

func Test(){
	test_1()
	test_2()
	test_3()
}

func test_1() {
	out := lengthOfLongestSubstring("abcadcbb")
	fmt.Println(out)
}
func test_2() {
	out := lengthOfLongestSubstring("pwwkew")
	fmt.Println(out)
}
func test_3() {
	out := lengthOfLongestSubstring("bbbbb")
	fmt.Println(out)
}

func lengthOfLongestSubstring(s string) int {

	m := make(map[uint8]int)
	longest := 0
	start_index := 0
	for i:=0; i < len(s) ; i++ {
		if old_index, exists := m[s[i]]; exists && old_index>=start_index {
			start_index = old_index + 1
		}
		m[s[i]] = i
		if length:=i-start_index+1; length > longest  {
			longest = length
		}
	}
	return longest
}
