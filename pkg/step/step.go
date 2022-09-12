package step

import (
	"tutero_assignment/pkg/src/graph"
)

type Stepper interface {
	// Step returns a prediction for the correct node; or an error if a prediction cannot be made.
	Step(graph graph.Graph) (graph.Node, error)
}

func New() *stepper {
	//* You may mutate this instantiation if necessary; but the function signature should not change.
	return &stepper{}
}

type stepper struct {
	ParentNodes []graph.Node
	ChildNodes  []graph.Node
	nodes       []graph.Node
	Check       bool
	Size        int
}

func (s *stepper) Step(graph graph.Graph) (graph.Node, error) {
	if s.nodes == nil {
		s.nodes = graph.Nodes()
		s.Size = len(graph.Nodes())
	}

	if s.Check == true {
		if s.Size == len(graph.Nodes()) {
			s.nodes = s.ChildNodes
		} else {
			s.nodes = s.ParentNodes
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
		s.Check = false
		s.Size = len(graph.Nodes())
	}
	s.nodes = graph.Nodes()
	if s.Size-1 > len(graph.Nodes()) {
		s.Size = len(graph.Nodes())
		s.Check = true
		if len(s.ParentNodes) > 0 {
			return s.ParentNodes[0], nil
		}
		return s.ChildNodes[0], nil
	}
	s.Size = len(graph.Nodes())
	node := s.nodes[s.Size/2]
	s.ParentNodes = graph.Parents(node)
	s.ChildNodes = graph.Children(node)
	return node, nil
}
