package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

type Node struct {
	parent      *Node
	children    []*Node
	name        string
	weight      int    // weight of this node itself
	totalweight int    // weight of this node and all subnodes
}

func PopulateSubtreeWeight(node *Node) int {
	weight := node.weight
	for _, child := range node.children {
		PopulateSubtreeWeight(child)
		weight += child.totalweight
	}
	node.totalweight = weight
	return weight
}

func GetUnbalancedChild(children []*Node) *Node {
	if len(children) == 0 {
		return nil
	}
	if len(children) < 3 {
		fmt.Errorf("Found node (%s) without enough children to decide which is correct\n", children[0].parent.name)
		os.Exit(1)
	}
	for pos, child := range children {
		if child.totalweight != children[(pos+1)%len(children)].totalweight &&
			child.totalweight != children[(pos+2)%len(children)].totalweight {
			return child
		}
	}
	return nil
}

func FindUnbalancedNode(node *Node, level int) *Node {
	var found *Node

	// Is the total subtree weight of any child unbalanced?
	child := GetUnbalancedChild(node.children)
	if child == nil {
		return nil
	}
	// If any grandchildren are unbalanced we need to go deeper in tree
	grandchild := GetUnbalancedChild(child.children)
	if grandchild != nil {
		found = FindUnbalancedNode(child, level+1)
	}
	// But if the grandchildren are balanced it means this child is unbalanced
	if found == nil {
		found = child
	}
	return found
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tree := make(map[string]*Node)
	var root *Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var node *Node
		for pos, token := range strings.Fields(scanner.Text()) {
			if token[len(token)-1] == ',' {
				token = token[:len(token)-1]
			}
			switch pos {
			case 0:
				if n, exist := tree[token]; exist {
					node = n
				} else {
					node = new(Node)
					tree[token] = node
					node.name = token
				}
				root = node // any starting point
			case 1:
				num, err := strconv.Atoi(token[1 : len(token)-1])
				if err != nil {
					panic(err)
				}
				node.weight = num
			case 2: // ignore ->
			default:
				var child *Node
				if n, exist := tree[token]; exist {
					child = n
				} else {
					child = new(Node)
					tree[token] = child
					child.name = token
				}
				child.parent = node
				node.children = append(node.children, child)
			}

		}
	}

	// root currently points to the last node added, find the real root node
	for {
		if root.parent == nil {
			break
		}
		root = root.parent
	}
	fmt.Printf("Root node: %s\n", root.name)

	PopulateSubtreeWeight(root)
	unb := FindUnbalancedNode(root, 0)
	if unb == nil {
		fmt.Printf("Could not find any unbalanced node.\n")
		os.Exit(1)
	}

	var sib *Node
	fmt.Printf("Unbalanced node: %s\n", unb.name)
	for _, child := range unb.children {
		fmt.Printf("  child -> %s (%d) (%d)\n", child.name, child.weight, child.totalweight)
	}
	for _, s := range unb.parent.children {
		if s.totalweight != unb.totalweight {
			sib = s
		}
		fmt.Printf("  sibling -> %s (%d) (%d)\n", s.name, s.weight, s.totalweight)
	}
	delta := sib.totalweight - unb.totalweight
	fmt.Printf("Node %s should weigh %d (%d%+d (+%d) = %d)\n",
		unb.name, unb.weight+delta, unb.weight, delta, unb.totalweight-unb.weight, sib.totalweight)

}
