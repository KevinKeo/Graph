/*
Package graph provides a generic framework to manipulate simple graph
Containing representation of different simple type of graphs and implements method to work with them
*/
package graph

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

var cptCoupleDirectedTest = 15
var testDirected_cases = []struct {
	matrix        [][]int
	nbArcs        int
	addArc        []couple
	removeArc     []couple
	directedGraph []IDirectedGraph
}{
	{
		[][]int{{math.MaxInt64, 1, math.MaxInt64, 1}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, 1}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64}},
		4,
		[]couple{{1, 3}, {0, 2}},
		[]couple{{3, 1}},
		[]IDirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, 1}, {math.MaxInt64, math.MaxInt64}},
		1,
		[]couple{},
		[]couple{},
		[]IDirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, 1, math.MaxInt64}, {1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64}},
		5,
		[]couple{},
		[]couple{},
		[]IDirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, math.MaxInt64, 1, 1, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64}},
		3,
		[]couple{},
		[]couple{},
		[]IDirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, math.MaxInt64, 1, 1, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1}, {1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}},
		7,
		[]couple{},
		[]couple{},
		[]IDirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, 1, 1, 1, 1}, {1, math.MaxInt64, 1, 1, 1}, {1, 1, math.MaxInt64, 1, 1}, {1, 1, 1, math.MaxInt64, 1}, {1, 1, 1, 1, math.MaxInt64}},
		20,
		[]couple{},
		[]couple{},
		[]IDirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}},
		0,
		[]couple{},
		[]couple{},
		[]IDirectedGraph{},
	},
}

func init() {
	for i := 0; i < len(testDirected_cases); i++ {
		matrix := testDirected_cases[i].matrix
		m1 := NewAdjacencyListDirectedGraphWithMatrix(matrix)
		m2 := NewAdjacencyMatrixDirectedGraphWithMatrix(matrix)
		m3 := NewAdjacencyListDirectedGraphWithInterface(m2)
		m4 := NewAdjacencyMatrixDirectedGraphWithInterface(m1)
		testDirected_cases[i].directedGraph = []IDirectedGraph{m1, m2, m3, m4}

		l := len(matrix)
		var s rand.Source
		var r *rand.Rand
		for cpt := 0; cpt < cptCoupleDirectedTest; cpt++ {
			s = rand.NewSource(time.Now().UnixNano() * time.Now().UnixNano() / int64((cpt+1)*(cpt+1)))
			r = rand.New(s)
			x := r.Intn(l)
			s = rand.NewSource(time.Now().UnixNano() * time.Now().UnixNano() * int64((cpt+100)*(cpt+100)))
			r = rand.New(s)
			y := r.Intn(l)
			testDirected_cases[i].addArc = append(testDirected_cases[i].addArc, couple{x, y})
			testDirected_cases[i].removeArc = append(testDirected_cases[i].removeArc, couple{x, y})
		}
	}
}

func TestInitDirectedGraph(t *testing.T) {
	for _, cases := range testDirected_cases {
		for _, graph := range cases.directedGraph {
			if graph.GetNbNodes() != len(cases.matrix) {
				t.Fatalf("Number of nodes should be equal to %d but is %d", len(cases.matrix), graph.GetNbNodes())
			}
			if graph.GetNbArcs() != cases.nbArcs {
				t.Fatalf("Number of arcs should be equal to %d but is %d", cases.nbArcs, graph.GetNbArcs())
			}
		}
	}
}

func TestAdjacencyMatrixDirectedGraph(t *testing.T) {
	for _, cases := range testDirected_cases {
		for _, graph := range cases.directedGraph {
			adjMatrix := graph.ToAdjacencyMatrix()
			for i, line := range adjMatrix {
				for j, v := range line {
					if cases.matrix[i][j] != v {
						t.Fatalf("The adjacent matrix is incorrect, the value at (%d,%d) should be equal to %d and not %d", i, j, cases.matrix[i][j], v)
					}
				}
			}
		}
	}
}

func TestComputeInverse(t *testing.T) {
	for _, cases := range testDirected_cases {
		for _, graph := range cases.directedGraph {
			invMatrix := graph.ComputeInverse().ToAdjacencyMatrix()
			for i, line := range invMatrix {
				for j, v := range line {
					if cases.matrix[j][i] != v {
						t.Fatalf("The adjacent matrix is incorrect, the value at (%d,%d) should be equal to %d and not %d", j, i, cases.matrix[j][i], v)
					}
				}
			}
		}
	}
}

func TestAddArc(t *testing.T) {
	for _, cases := range testDirected_cases {
		for _, coord := range cases.addArc {
			for _, graph := range cases.directedGraph {
				wasArc := graph.IsArc(coord.lig, coord.col)
				graph.AddArc(coord.lig, coord.col, 2)
				result := graph.IsArc(coord.lig, coord.col)
				if result {
					if coord.lig == coord.col {
						t.Fatalf("It should not be allowed to add an arc on the same node")
					} else if coord.lig < 0 || coord.col < 0 || coord.lig >= graph.GetNbNodes() || coord.col >= graph.GetNbNodes() {
						t.Fatalf("It should not be allowed to add an arc on non existant node")
					}
				} else {
					if coord.lig >= 0 && coord.col >= 0 && coord.lig < graph.GetNbNodes() && coord.col < graph.GetNbNodes() && coord.lig != coord.col {
						t.Fatalf("There should be an arc between from %d to %d.", coord.lig, coord.col)
					}
				}
				if !wasArc {
					graph.RemoveArc(coord.lig, coord.col)
				}
			}
		}

	}
}
func TestRemoveArc(t *testing.T) {
	for _, cases := range testDirected_cases {
		for _, coord := range cases.removeArc {
			for _, graph := range cases.directedGraph {
				wasArc := graph.IsArc(coord.lig, coord.col)
				graph.RemoveArc(coord.lig, coord.col)
				result := graph.IsArc(coord.lig, coord.col)
				if wasArc && result {
					t.Fatalf("The arc from %d to %d have not been removed by RemoveArc", coord.lig, coord.col)
				}
				if !wasArc && result {
					t.Fatalf("There was not an arc initialy from %d to %d and after using RemoveArc, there is", coord.lig, coord.col)
				}
				if wasArc {
					graph.AddArc(coord.lig, coord.col, 2)
				}
			}
		}

	}
}

func TestIsArc(t *testing.T) {
	for _, cases := range testDirected_cases {
		for _, graph := range cases.directedGraph {
			for i, col := range cases.matrix {
				for j, v := range col {
					isArc := graph.IsArc(i, j)
					if isArc && v == math.MaxInt64 {
						t.Fatalf("The method isArc result is %s but there should not have an arc from %d to %d.", isArc, i, j)
					}
					if !isArc && v != math.MaxInt64 {
						t.Fatalf("The method isArc result is %s but there should have an arc from %d to %d.", isArc, i, j)
					}
				}
			}
		}

	}
}

func TestGetSuccessors(t *testing.T) {
	for _, cases := range testDirected_cases {
		for _, graph := range cases.directedGraph {
			for from := 0; from < len(cases.matrix); from++ {
				for _, to := range graph.GetSuccessors(from) {
					if cases.matrix[from][to] == 0 {
						t.Fatalf("The node %d is not a successor of %d.", to, from)
					}
				}
			}
		}
	}
}

func TestGetPredecessors(t *testing.T) {
	for _, cases := range testDirected_cases {
		for _, graph := range cases.directedGraph {
			for to := 0; to < len(cases.matrix); to++ {
				for _, from := range graph.GetPredecessors(to) {
					if cases.matrix[from][to] == 0 {
						t.Fatalf("The node %d is not a predecessor of %d.", from, to)
					}
				}
			}
		}
	}
}

/*
func TestInitAdjacencyListInterface(t *testing.T) {
    for _, cases := range testDirected_cases {
        graphMatrix := NewAdjacencyListDirectedGraphWithMatrix(cases.matrix)
        graph := NewAdjacencyListDirectedGraphWithInterface(graphMatrix)
        if graphMatrix.NbArcs != graph.NbArcs {
            t.Fatalf("Number of arcs should be equal to %d but is %d", graphMatrix.NbArcs, graph.NbArcs)
        }
        if graphMatrix.NbNodes != graph.NbNodes {
            t.Fatalf("Number of Nodes should be equal to %d but is %d", graphMatrix.NbNodes, graph.NbNodes)
        }

        for i, v := range graphMatrix.listNode {
            if v != graph.listNode[i] {
                t.Fatalf("List Node in position %d should be equal to %d but is %d", i, v, graph.listNode[i])
            }
        }

        for i, v := range graphMatrix.succ {
            if v != graph.succ[i] {
                t.Fatalf("Successors in position %d should be equal to %d but is %d", i, v, graph.succ[i])
            }
        }
    }
}
*/
