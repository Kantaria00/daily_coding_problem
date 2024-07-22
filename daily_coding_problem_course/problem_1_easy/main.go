/*
Task:
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
Bonus:
Can you do this in one pass?
*/
package main

import (
	"fmt"
)

func main() {
	arr := []int{12, 10, 5, 7}
	k := 4
	fmt.Println(check(arr, k))
}
func check(arr []int, k int) bool {
	j := 1
	for i := 0; i < len(arr)-1; i++ {
		if i+j < len(arr) {
			switch {
			case arr[i]+arr[i+j] == k:
				return true
			case arr[i]+arr[i+j] != k:
				j++
				i--
			}
		} else {
			j = i + 2
		}
	}
	return false
}
