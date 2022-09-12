package step_test

import (
	"testing"
	g "tutero_assignment/pkg/src/graph"
	"tutero_assignment/pkg/step"
)

func TestStep(t *testing.T) {
	//* Implement unit-tests for your Step function
	graph := g.Graph{}
	graph.AddNode("A")
	graph.AddNode("B")
	stepper := step.New()
	submitted, _ := stepper.Step(graph)
	if submitted != graph.Nodes()[1] {
		t.Errorf("something wrong")
	}
}
