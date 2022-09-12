package step

import (
	g "tutero_assignment/pkg/src/graph"
)

type Stepper interface {
	// Step returns a prediction for the correct node; or an error if a prediction cannot be made.
	Step(graph g.Graph) (g.Node, error)
}

func New() *stepper {
	//* You may mutate this instantiation if necessary; but the function signature should not change.
	return &stepper{}
}

type stepper struct {
	parentNodes []g.Node
	childNodes  []g.Node
	nodes       []g.Node
	check       bool
	size        int
}

func (s *stepper) Step(graph g.Graph) (g.Node, error) {
	if s.nodes == nil {
		s.nodes = graph.Nodes()
		s.size = len(graph.Nodes())
	}

	if s.check == true {
		if s.size == len(graph.Nodes()) {
			s.nodes = s.childNodes
		} else {
			s.nodes = s.parentNodes
		}
		for _, itr := range graph.Nodes() {
			found := false
			for _, itr2 := range s.nodes {
				if itr == itr2 {
					found = true
				}
			}
			if found == false {
				graph.RemoveNode(itr)
			}
		}
		s.check = false
		s.size = len(graph.Nodes())
	}
	s.nodes = graph.Nodes()
	if s.size-1 > len(graph.Nodes()) {
		s.size = len(graph.Nodes())
		s.check = true
		if len(s.parentNodes) > 0 {
			return s.parentNodes[0], nil
		}
		return s.childNodes[0], nil
	}
	s.size = len(graph.Nodes())
	node := s.nodes[s.size/2]
	s.parentNodes = graph.Parents(node)
	s.childNodes = graph.Children(node)
	return node, nil
}
