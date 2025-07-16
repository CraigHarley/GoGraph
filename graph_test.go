package GoGraph

import "testing"

func setupTestGraph() *Node {
	h := NewNode("h", nil)
	f := NewNode("f", []*Node{h})
	e := NewNode("e", []*Node{f})
	d := NewNode("d", []*Node{e})
	c := NewNode("c", []*Node{d})
	b := NewNode("b", []*Node{c})
	a := NewNode("a", []*Node{b})

	// Add Cycles
	e.Edges = append(e.Edges, b) // e → b
	c.Edges = append(c.Edges, a) // c → a

	return a
}

type testCase struct {
	target   string
	expected bool
}

func TestDepthFirstSearch(t *testing.T) {
	root := setupTestGraph()

	tests := []testCase{
		{"a", true},
		{"h", true},
		{"z", false},
	}

	for _, tt := range tests {
		result := DepthFirstSearch(root, tt.target)
		if result != tt.expected {
			t.Errorf("DepthFirstSearch(%q) = %v; want %v", tt.target, result, tt.expected)
		}
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	root := setupTestGraph()

	tests := []testCase{
		{"a", true},
		{"h", true},
		{"z", false},
	}

	for _, tt := range tests {
		result := BreadthFirstSearch(root, tt.target)
		if result != tt.expected {
			t.Errorf("BreadthFirstSearch(%q) = %v; want %v", tt.target, result, tt.expected)
		}
	}
}
