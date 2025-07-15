package GoGraph

func NewNode(value string, edges []*Node) *Node {
	n := Node{
		Edges: edges,
		Value: value,
	}

	return &n
}

type Node struct {
	Edges []*Node
	Value string
}

// DepthFirstSearch
// todo - check for cycles (using visited map)
func DepthFirstSearch(from *Node, to string) bool {
	if from.Value == to {
		return true
	}
	for _, edge := range from.Edges {
		if DepthFirstSearch(edge, to) {
			return true
		}
	}
	return false
}

func BreadthFirstSearch(from *Node, to string) bool {
	visited := make(map[string]bool)
	queue := []*Node{from}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.Value] {
			continue
		}

		visited[current.Value] = true
		if current.Value == to {
			return true
		}

		queue = append(queue, current.Edges...)
	}
	return false
}
