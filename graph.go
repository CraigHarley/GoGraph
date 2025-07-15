package GoGraph

func NewNode(value string, edges []*Node) *Node {
	n := Node{
		edges: edges,
		value: value,
	}

	return &n
}

type Node struct {
	edges []*Node
	value string
}

func DepthFirstSearch(from *Node, to string) bool {
	if from.value == to {
		return true
	}
	for _, edge := range from.edges {
		return DepthFirstSearch(edge, to)
	}

	return false
}

func unshift(queue []*Node) ([]*Node, *Node) {
	current := queue[0]
	queue = queue[1:]

	return queue, current
}

func BreadthFirstSearch(from *Node, to string) bool {
	visited := make(map[string]bool)
	queue := []*Node{from}

	var current *Node
	for len(queue) > 0 {
		queue, current = unshift(queue)

		if visited[current.value] {
			continue
		}

		visited[current.value] = true
		if current.value == to {
			return true
		}

		for _, neighbor := range current.edges {
			queue = append(queue, neighbor)
		}
	}
	return false
}
