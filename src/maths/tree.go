package maths

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	inorder(root, res)
	return res
}

func inorder(root *TreeNode, res []int) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	inorder(root.Left, res)
	res = append(res, root.Val)
	inorder(root.Right, res)

}

func inorderTraversalBfs(root *TreeNode) []int {

	stack := make([]*TreeNode, 0)
	stack = append(stack, root.Left)
	result := make([]int, 0)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Left)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}

	return result

}
