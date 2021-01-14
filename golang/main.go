package main

import "fmt"

func main() {
	a := AddTwoNumbers(
		&ListNode{
			2,
			&ListNode{
				4,
				&ListNode{
					3,
					nil,
				},
			},
		},
		&ListNode{
			5,
			&ListNode{
				6,
				&ListNode{
					4,
					nil,
				},
			},
		},
	)
	fmt.Print("\nvalue:")
	for a.Next != nil {
		fmt.Print(a.Val)
		a = a.Next
	}
	fmt.Print(a.Val)
}
