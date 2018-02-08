package BST_nodes_in_range

type Node struct {
	value       int
	parent      *Node
	left, right *Node
}

func BuildBinarySearchTree(elements []int) *Node {
	var root *Node
	for _, element := range elements {
		root = Insert(root, element)
	}
	return root
}

func Insert(root *Node, element int) *Node {
	if root == nil {
		return &Node{value: element}
	}
	node := root
	for {
		if element < node.value {
			if node.left != nil {
				node = node.left
			} else {
				new := &Node{value: element, parent: node}
				node.left = new
				break
			}
		} else {
			if node.right != nil {
				node = node.right
			} else {
				new := &Node{value: element, parent: node}
				node.right = new
				break
			}
		}
	}
	return root
}

func GetNumbersInRange_recursive(start, end int, root *Node) []int {
	if root == nil {
		return []int{}
	}
	if root.value < start {
		return GetNumbersInRange_recursive(start, end, root.right)
	} else if root.value <= end {
		temp := append(GetNumbersInRange_recursive(start, end, root.left), root.value)
		return append(temp, GetNumbersInRange_recursive(start, end, root.right)...)
	} else {
		return GetNumbersInRange_recursive(start, end, root.left)
	}
}
