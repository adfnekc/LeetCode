package main

import "fmt"

// Print is
var (
	P  = fmt.Print
	Pl = fmt.Println
	Pf = fmt.Printf
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a, b := []int{1}, []int{2, 4}
	Pf("%f", F(a, b))
}
