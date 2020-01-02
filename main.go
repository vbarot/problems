package main

import (
	"fmt"
	"problems/ps/add_two_numbers"
	"problems/ps/longest_uniquechar_substring"
	"problems/ps/two_sum"
	"os"
)

func main(){

	if len(os.Args) != 2 {
		fmt.Println("Invalid input, problems takes exactly 1 argument")
		os.Exit(-1)
	}

	switch os.Args[1] {
	case "two_sum":
		two_sum.Describe()
		two_sum.Test()
	case "add_two_numbers":
		add_two_numbers.Describe()
		add_two_numbers.Test()
	case "longest_uniquechar_substring":
		longest_uniquechar_substring.Describe()
		longest_uniquechar_substring.Test()
	default:
		fmt.Println("Invalid Input")
		os.Exit(-2)
	}



}
