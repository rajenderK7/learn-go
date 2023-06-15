package lstruct

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
	}
	if val <= root.val {
		root.left = insertNode(root.left, val)
	} else {
		root.right = insertNode(root.right, val)
	}
	return root
}

func (root *TreeNode) insertNode(val int) {
	if val <= root.val {
		if root.left != nil {
			root.left.insertNode(val)
		} else {
			root.left = NewTreeNode(val)
		}
	} else {
		if root.right != nil {
			root.right.insertNode(val)
		} else {
			root.right = NewTreeNode(val)
		}
	}
}

func (root *TreeNode) PrintInorder() {
	if root != nil {
		root.left.PrintInorder()
		fmt.Printf("%d ", root.val)
		root.right.PrintInorder()
	}
}

func ConstrucBSTfromSlice(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// var root *TreeNode
	root := NewTreeNode(nums[0])
	for i := 1; i < len(nums); i++ {
		root.insertNode(nums[i])
	}
	// for _, num := range nums {
	// 	root = insertNode(root, num)
	// }
	return root
}

func (root *TreeNode) inorderPredecessor() *TreeNode {
	ptr := root.left
	for ptr.right != nil && ptr.right != root {
		ptr = ptr.right
	}
	return ptr
}

func (root *TreeNode) MorrisInorderTraversal() []int {
	res := []int{}
	ptr := root
	for ptr != nil {
		if ptr.left == nil {
			res = append(res, ptr.val)
			ptr = ptr.right
		} else {
			pred := ptr.inorderPredecessor()
			if pred == nil {
				pred.right = ptr
				ptr = ptr.left
			} else {
				res = append(res, ptr.val)
				pred.right = nil
				ptr = ptr.right
			}
		}
	}
	return res
}
