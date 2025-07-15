package GoGraph

import "testing"

func setupTestGraph() *Node {
	h := NewNode("h", nil)
	f := NewNode("f", []*Node{h})
	e := NewNode("e", []*Node{f})
	d := NewNode("d", []*Node{e})
	c := NewNode("c", []*Node{d})
	b := NewNode("b", []*Node{c})
	a := NewNode("a", []*Node{b, h})

	return a
}

func TestDepthFirstSearch(t *testing.T) {
	root := setupTestGraph()

	tests := []struct {
		target   string
		expected bool
	}{
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

	tests := []struct {
		target   string
		expected bool
	}{
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
