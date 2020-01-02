package add_two_numbers

import (
	"fmt"
	"problems/ds/list"
)

type ListNode list.ListNode

func Describe() {
	const desc = `##########
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
	Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Output: 7 -> 0 -> 8
	Explanation: 342 + 465 = 807.
##########`
	fmt.Println(desc)
}

func Test(){
	test_1()
}

func test_1() {
	l1 := &list.List{}
	l1.Add(5)
	l2 := &list.List{}
	l2.Add(5)
	out := addTwoNumbers((*ListNode)(l1.Root), (*ListNode)(l2.Root))
	curr_out := out
	for curr_out != nil {
		fmt.Print(curr_out.Val)
		curr_out = (*ListNode)(curr_out.Next)
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	output := &ListNode{}
	current_num_1 :=l1
	current_num_2 :=l2
	current_output := output

	carry := 0
	for current_num_1 != nil || current_num_2 != nil || carry > 0 {
		digit_1, digit_2 := 0,0
		if current_num_1 != nil{
			digit_1 = current_num_1.Val
			current_num_1 = (*ListNode)(current_num_1.Next)
		}
		if current_num_2 != nil{
			digit_2 = current_num_2.Val
			current_num_2 = (*ListNode)(current_num_2.Next)
		}
		output_val := digit_1 + digit_2 + carry
		current_output.Val = output_val%10
		carry = output_val/10
		if current_num_1 != nil || current_num_2 != nil || carry > 0 {
			current_output.Next = (*list.ListNode)((&ListNode{}))
			current_output = (*ListNode)(current_output.Next)
		}
	}
	return output
}
