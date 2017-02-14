/*
Package graph provides a generic framework to manipulate simple graph
Containing representation of different simple type of graphs and implements method to work with them
*/
package graph

import "math"

/*
AdjacencyMatrixUndirectedGraph represent a Undirected Adjacency Matrix
*/
type AdjacencyMatrixUndirectedGraph struct {
	NbNodes, NbEdges int
	matrice          [][]int
}

/*
NewAdjacencyMatrixUndirectedGraph initialize a pointer to a new AdjacencyMatrixUndirectedGraph with a [][]int parameter
*/
func NewAdjacencyMatrixUndirectedGraphWithMatrix(mat [][]int) *AdjacencyMatrixUndirectedGraph {
	var edges int
	nodes := len(mat)
	for i := 0; i < nodes; i++ {
		for j := 0; j < nodes; j++ {
			if mat[i][j] != math.MaxInt64 {
				edges++
			}
		}
	}
	return &AdjacencyMatrixUndirectedGraph{nodes, edges / 2, mat}
}

//NewAdjacencyMatrixUndirectedGraphWithInterface initialize a pointer to a new AdjacencyMatrixUndirectedGraph taking in parameter a IUndirectedGraph
func NewAdjacencyMatrixUndirectedGraphWithInterface(undirectedGraph IUndirectedGraph) *AdjacencyMatrixUndirectedGraph {
	return &AdjacencyMatrixUndirectedGraph{undirectedGraph.GetNbNodes(), undirectedGraph.GetNbEdges(), undirectedGraph.ToAdjacencyMatrix()}
}

//GetNbEdges gives the number of edges in the graph
func (a AdjacencyMatrixUndirectedGraph) GetNbEdges() (n int) {
	return a.NbEdges
}

//GetNbNodes gives the number of nodes in the graph
func (a AdjacencyMatrixUndirectedGraph) GetNbNodes() (n int) {
	return a.NbNodes
}

//ToAdjacencyMatrix return the adjacency matrix
func (a AdjacencyMatrixUndirectedGraph) ToAdjacencyMatrix() [][]int {
	return a.matrice
}

//IsEdge return true if there is an edge between i and j
func (a AdjacencyMatrixUndirectedGraph) IsEdge(i int, j int) bool {
	if i >= len(a.matrice) || j >= len(a.matrice) || i < 0 || j < 0 {
		return false
	}
	return a.matrice[i][j] != math.MaxInt64 && a.matrice[j][i] == a.matrice[i][j]
}

//AddEdge add a new edge between i and j, requires i != j
func (a *AdjacencyMatrixUndirectedGraph) AddEdge(i int, j int, p int) {
	if i >= len(a.matrice) || j >= len(a.matrice) || i == j || i < 0 || j < 0 {
		return
	}
	if a.matrice[i][j] != math.MaxInt64 {
		return
	}
	a.matrice[i][j] = p
	a.matrice[j][i] = p
	a.NbEdges += 1
}

//RemoveEdge remove a edge between i and j, requires i!=j
func (a *AdjacencyMatrixUndirectedGraph) RemoveEdge(i int, j int) {
	if i >= len(a.matrice) || j >= len(a.matrice) || i == j || i < 0 || j < 0 {
		return
	}
	if a.matrice[i][j] == math.MaxInt64 {
		return
	}
	a.matrice[i][j] = math.MaxInt64
	a.matrice[j][i] = math.MaxInt64
	a.NbEdges -= 1
	/*
	   a.Matrice = GenerateGraphData(1, 0, false)
	   a.Matrice[i][j] = 7
	   a.Matrice[j][i] = 7*/

}

//GetNeighbors returns a new slice of int containing neighbors of node i
func (a AdjacencyMatrixUndirectedGraph) GetNeighbors(i int) (neighbors []int) {
	for node, value := range a.matrice[i] {
		if value != math.MaxInt64 {
			neighbors = append(neighbors, node)
		}
	}
	return neighbors
}

func (a AdjacencyMatrixUndirectedGraph) GetWeight(x, y int) int {
	return a.matrice[x][y]
}
