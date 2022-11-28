package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

const (
	BLUE   = "\033[1;34m"
	GRN    = "\033[1;32m"
	RED    = "\033[0;31m"
	VIOLET = "\033[0;35m"
	YELLOW = "\033[1;33m"
	TICK   = "\xE2\x9C\x94"
	END    = "\033[0m"
)

func countToys(tree *TreeNode, sum int) int {
	if tree.Left != nil {
		sum = countToys(tree.Left, sum)
	}
	if tree.Right != nil {
		sum = countToys(tree.Right, sum)
	}
	if tree.HasToy {
		sum++
	}
	return sum
}

func areToysBalanced(tree *TreeNode) bool {
	return countToys(tree.Left, 0) == countToys(tree.Right, 0)
}

func printTree(prefix string, tree *TreeNode, isLeft bool, level int) {
	if tree != nil {
		fmt.Print(level, prefix)
		if isLeft {
			fmt.Print("├──")
		} else {
			fmt.Print("└──")
		}
		if tree.HasToy {
			fmt.Println(GRN, 1, END)
		} else {
			fmt.Println(VIOLET, 0, END)
		}
		if isLeft {
			printTree(prefix+"│   ", tree.Left, true, level+1)
		} else {
			printTree(prefix+"    ", tree.Left, true, level+1)
		}
		if isLeft {
			printTree(prefix+"│   ", tree.Right, false, level+1)
		} else {
			printTree(prefix+"    ", tree.Right, false, level+1)
		}
	}
}

func main() {
	var tree1 = &TreeNode{false,
		&TreeNode{false,
			&TreeNode{false, nil, nil},
			&TreeNode{true, nil, nil}},
		&TreeNode{true, nil, nil}}
	printTree("", tree1, false, 0)
	fmt.Println(YELLOW, "areToysBalanced", END, areToysBalanced(tree1))

	var tree2 = &TreeNode{true,
		&TreeNode{true,
			&TreeNode{true, nil, nil},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}

	printTree("", tree2, false, 0)
	fmt.Println(YELLOW, "areToysBalanced", END, areToysBalanced(tree2))

	var tree3 = &TreeNode{true,
		&TreeNode{true, nil, nil},
		&TreeNode{false, nil, nil}}
	printTree("", tree3, false, 0)
	fmt.Println(YELLOW, "areToysBalanced", END, areToysBalanced(tree3))

	var tree4 = &TreeNode{false,
		&TreeNode{true, nil,
			&TreeNode{true, nil, nil}},
		&TreeNode{false, nil,
			&TreeNode{true, nil, nil}}}
	printTree("", tree4, false, 0)
	fmt.Println(YELLOW, "areToysBalanced", END, areToysBalanced(tree4))
}
