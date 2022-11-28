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

func goThroughLevel(level int, curLevel int, tree *TreeNode) ([]bool, error) {
	var result []bool
	var tmp []bool
	if level%2 == 0 {
		if tree.Right != nil && curLevel != level {
			tmp, _ := goThroughLevel(level, curLevel+1, tree.Right)
			result = append(result, tmp...)
		}
		if tree.Left != nil && curLevel != level {
			tmp, _ := goThroughLevel(level, curLevel+1, tree.Left)
			result = append(result, tmp...)
		}
	} else {
		if tree.Left != nil && curLevel != level {
			tmp, _ := goThroughLevel(level, curLevel+1, tree.Left)
			result = append(result, tmp...)
		}
		if tree.Right != nil && curLevel != level {
			tmp, _ := goThroughLevel(level, curLevel+1, tree.Right)
			result = append(result, tmp...)
		}
	}
	if curLevel == level {
		result = append(result, tree.HasToy)
	}
	_ = tmp
	return result, nil
}

func unrollGarland(tree *TreeNode) []bool {
	var result []bool
	var tmp []bool
	for i := 0; i < 10; i++ {
		tmp, _ = goThroughLevel(i, 0, tree)
		result = append(result, tmp...)
	}
	return result
}

func main() {
	var tree1 = &TreeNode{false,
		&TreeNode{false,
			&TreeNode{false, nil, nil},
			&TreeNode{true, nil, nil}},
		&TreeNode{true, nil, nil}}
	printTree("", tree1, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree1))

	var tree2 = &TreeNode{true,
		&TreeNode{true,
			&TreeNode{true, nil, nil},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}

	printTree("", tree2, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree2))

	var tree3 = &TreeNode{true,
		&TreeNode{true, nil, nil},
		&TreeNode{false, nil, nil}}
	printTree("", tree3, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree3))

	var tree4 = &TreeNode{false,
		&TreeNode{true, nil,
			&TreeNode{true, nil, nil}},
		&TreeNode{false, nil,
			&TreeNode{true, nil, nil}}}
	printTree("", tree4, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree4))

	var tree5 = &TreeNode{true, nil, nil}
	printTree("", tree5, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree5))

	var tree6 = &TreeNode{false, nil, nil}
	printTree("", tree6, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree6))

	var tree7 = &TreeNode{false, nil, &TreeNode{false, nil, nil}}
	printTree("", tree7, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree7))

	var tree8 = &TreeNode{false, nil, &TreeNode{true, nil, nil}}
	printTree("", tree8, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree8))

	var tree9 = &TreeNode{true,
		&TreeNode{true,
			&TreeNode{true, nil, &TreeNode{true,
				&TreeNode{true,
					&TreeNode{true, nil, nil},
					&TreeNode{false, nil, nil}},
				&TreeNode{false,
					&TreeNode{true, nil, nil},
					&TreeNode{true, nil, nil}}}},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}
	printTree("", tree9, false, 0)
	fmt.Println(YELLOW, "unrollGarland", END, unrollGarland(tree9))
}
