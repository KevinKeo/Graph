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

var cptCoupleUndirectedTest = 10
var testUndirected_cases = []struct {
	matrix          [][]int
	nbEdges         int
	addEdge         []couple
	removeEdge      []couple
	undirectedGraph []IUndirectedGraph
}{
	{
		[][]int{{math.MaxInt64, 1, 1, 1}, {1, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {1, math.MaxInt64, math.MaxInt64, 1}, {1, math.MaxInt64, 1, math.MaxInt64}},
		4,
		[]couple{{1, 2}, {1, 3}, {3, 3}, {2, 3}},
		[]couple{{1, 3}, {1, 1}, {2, 3}, {1, 2}},
		[]IUndirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, 1}, {1, math.MaxInt64, 1, 1, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, 1}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {1, math.MaxInt64, 1, math.MaxInt64, math.MaxInt64}},
		5,
		[]couple{},
		[]couple{},
		[]IUndirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, 1, math.MaxInt64, 1, 1}, {1, math.MaxInt64, 1, 1, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, 1, 1}, {1, 1, 1, math.MaxInt64, 1}, {1, math.MaxInt64, 1, 1, math.MaxInt64}},
		8,
		[]couple{},
		[]couple{},
		[]IUndirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, 1, 1, 1, 1}, {1, math.MaxInt64, 1, 1, 1}, {1, 1, math.MaxInt64, 1, 1}, {1, 1, 1, math.MaxInt64, 1}, {1, 1, 1, 1, math.MaxInt64}},
		10,
		[]couple{},
		[]couple{},
		[]IUndirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, 1, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {1, math.MaxInt64, math.MaxInt64, 1, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1, 1, 1, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}},
		8,
		[]couple{},
		[]couple{},
		[]IUndirectedGraph{},
	},
	{
		[][]int{{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}},
		0,
		[]couple{},
		[]couple{},
		[]IUndirectedGraph{},
	},
}

func init() {
	for i := 0; i < len(testUndirected_cases); i++ {
		matrix := testUndirected_cases[i].matrix
		m1 := NewAdjacencyMatrixUndirectedGraphWithMatrix(matrix)
		m2 := NewAdjacencyListUndirectedGraphWithMatrix(matrix)
		m3 := NewAdjacencyMatrixUndirectedGraphWithInterface(m2)
		m4 := NewAdjacencyListUnirectedGraphWithInterface(m1)
		testUndirected_cases[i].undirectedGraph = []IUndirectedGraph{m1, m2, m3, m4}

		l := len(matrix)
		var s rand.Source
		var r *rand.Rand
		for cpt := 0; cpt < cptCoupleUndirectedTest; cpt++ {
			s = rand.NewSource(time.Now().UnixNano() * time.Now().UnixNano() / int64((cpt+1)*(cpt+1)))
			r = rand.New(s)
			x := r.Intn(l)
			s = rand.NewSource(time.Now().UnixNano() * time.Now().UnixNano() * int64((cpt+100)*(cpt+100)))
			r = rand.New(s)
			y := r.Intn(l)
			testUndirected_cases[i].addEdge = append(testUndirected_cases[i].addEdge, couple{x, y})
			testUndirected_cases[i].removeEdge = append(testUndirected_cases[i].removeEdge, couple{x, y})
		}
	}
}

/* dont WORK problem pointer ???
func init() {
	for _, cases := range testUndirected_cases {
		cases.undirectedGraph = []IUndirectedGraph{NewAdjacencyListUndirectedGraphWithMatrix(cases.matrix)}
	}
}
*/
func TestInitAdjacencyList(t *testing.T) {
	for _, cases := range testUndirected_cases {
		for _, graph := range cases.undirectedGraph {
			if graph.GetNbNodes() != len(cases.matrix) {
				t.Fatalf("Number of nodes should be equal to %d but is %d", len(cases.matrix), graph.GetNbNodes())
			}
			if graph.GetNbEdges() != cases.nbEdges {
				t.Fatalf("Number of edges should be equal to %d but is %d", cases.nbEdges, graph.GetNbEdges())
			}
		}
	}
}

func TestAdjacencyMatrix(t *testing.T) {
	for _, cases := range testUndirected_cases {
		for _, graph := range cases.undirectedGraph {
			adjMatrix := graph.ToAdjacencyMatrix()
			for i, line := range adjMatrix {
				for j, v := range line {
					if cases.matrix[i][j] != v {
						t.Fatalf("The adjacent matrix is incorrect, the value at (%d,%d) should be equal to %d and not %d", i, j, cases.matrix[i][j], v)
					}
					if cases.matrix[i][j] != cases.matrix[j][i] {
						t.Fatalf("The adjacent matrix for a undirected graph should be symetric, or the value at (%d,%d) is equal to %d and for (%d,%d) is %d", i, j, cases.matrix[i][j], j, i, cases.matrix[i][j])
					}
				}
			}
		}
	}
}

func TestAddEdge(t *testing.T) {
	for _, cases := range testUndirected_cases {
		for _, coord := range cases.addEdge {
			for _, graph := range cases.undirectedGraph {
				wasEdge := graph.IsEdge(coord.lig, coord.col)
				nbEdge := graph.GetNbEdges()
				graph.AddEdge(coord.lig, coord.col, 2)
				isEdge := graph.IsEdge(coord.lig, coord.col)
				if isEdge {
					if coord.lig == coord.col {
						t.Fatalf("It should not be allowed to add an edge on the same node (%d,%d)", coord.lig, coord.col)
					} else if coord.lig < 0 || coord.col < 0 || coord.lig >= graph.GetNbNodes() || coord.col >= graph.GetNbNodes() {
						t.Fatalf("It should not be allowed to add an edge on non existant node (%d,%d)", coord.lig, coord.col)
					}
					if graph.GetNbEdges() != nbEdge+1 && !wasEdge {
						t.Fatalf("The edge (%d,%d) was add, but the number of edge was %d, and now is %d and it should be %d", coord.lig, coord.col, nbEdge, graph.GetNbEdges(), nbEdge+1)
					}
					if graph.GetNbEdges() != nbEdge && wasEdge {
						t.Fatalf("The edge (%d,%d) already existed, but the number of edge was %d, and now is %d and it should be %d", coord.lig, coord.col, nbEdge, graph.GetNbEdges(), nbEdge)
					}
				} else {
					if coord.lig >= 0 && coord.col >= 0 && coord.lig < graph.GetNbNodes() && coord.col < graph.GetNbNodes() && coord.lig != coord.col {
						t.Fatalf("There should be an edge between %d and %d.", coord.lig, coord.col, graph)
					}
					if graph.GetNbEdges() != nbEdge {
						t.Fatalf("There was not edge (%d,%d) but the number of edge was modified from %d to %d", coord.lig, coord.col, nbEdge, graph.GetNbEdges())
					}
				}
				if !wasEdge {
					graph.RemoveEdge(coord.lig, coord.col)
				}
			}
		}

	}
}

func TestRemoveEdge(t *testing.T) {
	for _, cases := range testUndirected_cases {
		for _, coord := range cases.removeEdge {
			for _, graph := range cases.undirectedGraph {
				wasEdge := graph.IsEdge(coord.lig, coord.col)
				nbEdge := graph.GetNbEdges()
				graph.RemoveEdge(coord.lig, coord.col)
				isEdge := graph.IsEdge(coord.lig, coord.col)
				if wasEdge {
					if isEdge {
						t.Fatalf("The edge between %d and %d have not been removed by RemoveEdge", coord.lig, coord.col)
					}
					if graph.GetNbEdges() != nbEdge-1 {
						t.Fatalf("The edge (%d,%d) was removed, but the number of edge was %d, and now is %d and it should be %d", coord.lig, coord.col, nbEdge, graph.GetNbEdges(), nbEdge-1)
					}
					graph.AddEdge(coord.lig, coord.col, 2)
				} else {
					if isEdge {
						t.Fatalf("There was not an edge initialy between %d and %d and after using RemoveEdge, there is", coord.lig, coord.col)
					}
					if graph.GetNbEdges() != nbEdge {
						t.Fatalf("There was not edge (%d,%d), but the number of edge was modified from %d to %d", coord.lig, coord.col, nbEdge, graph.GetNbEdges())
					}
				}
			}
		}

	}
}

func TestIsEdge(t *testing.T) {
	for _, cases := range testUndirected_cases {
		for _, graph := range cases.undirectedGraph {
			for i, col := range cases.matrix {
				for j, v := range col {
					isEdge := graph.IsEdge(i, j)
					if isEdge && v == math.MaxInt64 {
						t.Fatalf("The method IsEdge result is %s but there should not have an edge between %d and %d.", isEdge, i, j)
					}
					if !isEdge && v != math.MaxInt64 {
						t.Fatalf("The method IsEdge result is %s but there should have an edge between %d and %d.", isEdge, i, j)
					}
				}
			}
		}

	}
}

func TestGetNeighbors(t *testing.T) {
	for _, cases := range testUndirected_cases {
		for _, graph := range cases.undirectedGraph {
			for i := 0; i < len(cases.matrix); i++ {
				cpt := 0
				for j := 0; j < len(cases.matrix[i]); j++ {
					if cases.matrix[i][j] != math.MaxInt64 {
						cpt++
					}
				}
				neighbors := graph.GetNeighbors(i)
				if len(neighbors) != cpt {
					t.Fatalf("The node %d should have %d neighbors but have %d", i, len(neighbors), cpt)
				}
				for _, j := range graph.GetNeighbors(i) {
					if cases.matrix[i][j] == math.MaxInt64 {
						t.Fatalf("The node %d should have been a neighbor of %d.", i, j)
					}
				}
			}
		}
	}
}
