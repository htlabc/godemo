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

func inorderTraversalBfs_before(root *TreeNode) []int {

	stack := make([]*TreeNode, 0)
	stack = append(stack, root.Left)
	result := make([]int, 0)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}

	return result

}

func inorderTraversalBfs_middle(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}

	return result
}

//通过的迭代获取 树的前序查询
func inorderTraversalBfs_middle2(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) != 0 {
		//获取栈顶节点
		node := stack[len(stack)-1]
		if node != nil {
			stack = stack[:len(stack)-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			//如果头部节点为空，则弹出栈顶数据并且加入到结果数据集
		} else {
			stack = stack[:len(stack)-1]
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
		}
	}

	return result

}
