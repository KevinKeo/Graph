/*
Package graph provides a generic framework to manipulate simple graph
Containing representation of different simple type of graphs and implements method to work with them
*/
package graph

import (
	"testing"
)

type couple struct {
	lig, col int
}

func TestGenerateGraphSym(t *testing.T) {
	length := 8
	edge := 5
	graph := GenerateGraphData(length, edge, true)
	if len(graph) != length {
		t.Fatalf("Length should be equal to %d but is %d", length, len(graph))
	}
	cptEdge := 0
	for i, v1 := range graph {
		for j := range v1 {
			if i < j {
				break
			}
			if i == j {
				if graph[i][j] != 0 {
					t.Fatalf("The diagonal of the matrix should be equal to 0 and not %d", graph[i][j])
				}
			}
			if graph[i][j] != graph[j][i] {
				t.Fatalf("The value in (%d,%d)=%d should be equal to (%d,%d)=%d", i, j, graph[i][j], j, i, graph[j][i])
			}
			if graph[i][j] != 0 {
				cptEdge++
			}
		}
	}
	if cptEdge != edge {
		t.Fatalf("Number of edges should be equal to %d but is %d", edge, cptEdge)
	}
}

func TestGenerateGraph(t *testing.T) {
	length := 8
	edge := 9
	graph := GenerateGraphData(length, edge, false)
	if len(graph) != length {
		t.Fatalf("Length should be equal to %d but is %d", length, len(graph))
	}
	cptEdge := 0
	for i, v1 := range graph {
		for j := range v1 {
			if i == j {
				if graph[i][j] != 0 {
					t.Fatalf("The diagonal of the matrix should be equal to 0 and not %d", graph[i][j])
				}
			}
			if graph[i][j] != 0 {
				cptEdge++
			}
		}
	}
	if cptEdge != edge {
		t.Fatalf("Number of edges should be equal to %d but is %d", edge, cptEdge)
	}
}
