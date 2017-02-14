/*
Package graph provides a generic framework to manipulate simple graph
Containing representation of different simple type of graphs and implements method to work with them
*/
package graph

import (
	"math"
)

/*
 */
type AdjacencyMatrixDirectedGraph struct {
	NbNodes, nbArcs int
	matrice         [][]int
}

/*
NewAdjacencyMatrixDirectedGraphWithMatrix initialize a pointer to a new AdjacencyMatrixDirectedGraph with a [][]int parameter
*/
func NewAdjacencyMatrixDirectedGraphWithMatrix(mat [][]int) *AdjacencyMatrixDirectedGraph {
	arcs := 0
	nodes := len(mat)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if mat[i][j] != math.MaxInt64 {
				arcs++
			}
		}
	}
	return &AdjacencyMatrixDirectedGraph{nodes, arcs, mat}
}

//NewAdjacencyMatrixDirectedGraphWithInterface initialize a pointer to a new AdjacencyMatrixDirectedGraph taking in parameter a IDirectedGraph
func NewAdjacencyMatrixDirectedGraphWithInterface(graph IDirectedGraph) *AdjacencyMatrixDirectedGraph {
	return &AdjacencyMatrixDirectedGraph{graph.GetNbNodes(), graph.GetNbArcs(), graph.ToAdjacencyMatrix()}
}

//GetNbNodes returns the number of nodes of the adjacencyMatrixDirectedGraph
func (a AdjacencyMatrixDirectedGraph) GetNbNodes() int {
	return a.NbNodes
}

//ToAdjacencyMatrix return the adjacency matrix of the graph
func (a AdjacencyMatrixDirectedGraph) ToAdjacencyMatrix() [][]int {
	return a.matrice
}

func (a AdjacencyMatrixDirectedGraph) GetNbArcs() int {
	return a.nbArcs
}

func (a AdjacencyMatrixDirectedGraph) IsArc(x, y int) bool {
	if x < 0 || y < 0 || x >= a.NbNodes || y >= a.NbNodes {
		return false
	}
	if a.matrice[x][y] != math.MaxInt64 {
		return true
	}
	return false
}

func (a *AdjacencyMatrixDirectedGraph) RemoveArc(x, y int) {
	if x < 0 || y < 0 || x >= a.NbNodes || y >= a.NbNodes {
		return
	}
	if a.matrice[x][y] == math.MaxInt64 {
		return
	}
	a.matrice[x][y] = math.MaxInt64
	a.nbArcs--
}

func (a *AdjacencyMatrixDirectedGraph) AddArc(x, y, p int) {
	if x < 0 || y < 0 || x >= a.NbNodes || y >= a.NbNodes || x == y {
		return
	}
	if a.matrice[x][y] != math.MaxInt64 {
		return
	}
	a.matrice[x][y] = p
	a.nbArcs++
}

func (a AdjacencyMatrixDirectedGraph) GetSuccessors(x int) (succ []int) {
	if x < 0 && x >= a.NbNodes {
		return succ
	}
	for n, v := range a.matrice[x] {
		if v != math.MaxInt64 {
			succ = append(succ, n)
		}
	}
	return succ
}

func (a AdjacencyMatrixDirectedGraph) GetPredecessors(x int) (pred []int) {
	if x < 0 && x >= a.NbNodes {
		return pred
	}
	for i := 0; i < a.NbNodes; i++ {
		if a.matrice[i][x] != math.MaxInt64 {
			pred = append(pred, i)
		}
	}
	return pred
}

func (a AdjacencyMatrixDirectedGraph) ComputeInverse() IDirectedGraph {
	mat := make([][]int, a.NbNodes)
	for n := range a.matrice {
		mat[n] = make([]int, a.NbNodes)
	}

	for i := 0; i < a.NbNodes; i++ {
		for j := 0; j < a.NbNodes; j++ {
			mat[j][i] = a.matrice[i][j]
		}
	}
	return &AdjacencyMatrixDirectedGraph{a.NbNodes, a.NbNodes, mat}
}

func (a AdjacencyMatrixDirectedGraph) GetWeight(x, y int) int {
	return a.matrice[x][y]
}
