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

func DepthFirstSearch(from *Node, to string) bool {
	visited := make(map[string]bool)

	return dfs(from, to, visited)
}

func dfs(current *Node, target string, visited map[string]bool) bool {
	if current.Value == target {
		return true
	}

	if visited[current.Value] {
		return false
	}

	visited[current.Value] = true

	for _, edge := range current.Edges {
		if dfs(edge, target, visited) {
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
