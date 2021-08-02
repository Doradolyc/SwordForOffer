package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间复杂度 o(n*n)
// 空间复杂度 o(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	mid := preorder[0]

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == mid {
			mid = i
			break
		}
	}

	root := &TreeNode{
		Val: preorder[0],
		Left: buildTree(preorder[1 : 1 + mid], inorder[0 : mid]),
		Right: buildTree(preorder[mid + 1 : len(preorder)], inorder[mid + 1 : len(inorder)]),
	}

	return root
}