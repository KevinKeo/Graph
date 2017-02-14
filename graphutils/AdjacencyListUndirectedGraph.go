/*
Package graph provides a generic framework to manipulate simple graph
Containing representation of different simple type of graphs and implements method to work with them
*/
package graph

import "math"

//AdjacencyListUndirectedGraph represents an undirected graph in the form of an adjacency List
type AdjacencyListUndirectedGraph struct {
	NbNodes, NbEdges       int
	listNode, succ, weight []int
}

//NewAdjacencyListUndirectedGraph create a new AdjacencyListUndirectedGraph's pointer
func NewAdjacencyListUndirectedGraphWithMatrix(generatedGraph [][]int) *AdjacencyListUndirectedGraph {
	nodes := make([]int, len(generatedGraph)+1)
	var succ, weight []int
	for i, s := range generatedGraph {
		for j, v := range s {
			if v != math.MaxInt64 {
				succ = append(succ, j)
				weight = append(weight, v)
			}
		}
		nodes[i+1] = len(succ)
	}
	return &AdjacencyListUndirectedGraph{len(nodes) - 1, len(succ) / 2, nodes, succ, weight}
}

//NewAdjacencyListUnirectedGraphWithInterface create a new AdjacencyListUndirectGraph's pointer
func NewAdjacencyListUnirectedGraphWithInterface(undirectedGraph IUndirectedGraph) *AdjacencyListUndirectedGraph {
	var succ []int
	var weight []int
	nodes := make([]int, undirectedGraph.GetNbNodes()+1)
	for i := 0; i < undirectedGraph.GetNbNodes(); i++ {
		succ = append(succ, undirectedGraph.GetNeighbors(i)...)
		nodes[i+1] = len(succ)
		for j := nodes[i]; j < nodes[i+1]; j++ {
			weight = append(weight, undirectedGraph.GetWeight(i, succ[j]))
		}
	}
	return &AdjacencyListUndirectedGraph{undirectedGraph.GetNbNodes(), undirectedGraph.GetNbEdges(), nodes, succ, weight}
}

//GetNbNodes return the number of Nodes of the graph
func (a AdjacencyListUndirectedGraph) GetNbNodes() int {
	return a.NbNodes
}

func (a AdjacencyListUndirectedGraph) ToAdjacencyMatrix() [][]int {
	matrix := make([][]int, a.GetNbNodes())
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, a.NbNodes)
		for j := 0; j < len(matrix); j++ {
			matrix[i][j] = math.MaxInt64
		}
	}
	for i := 0; i < a.GetNbNodes(); i++ {
		for j := a.listNode[i]; j < a.listNode[i+1]; j++ {
			matrix[i][a.succ[j]] = a.weight[j]
		}
	}
	return matrix
}

//GetNbEdges gives the number of edges in the graph
func (a AdjacencyListUndirectedGraph) GetNbEdges() int {
	return a.NbEdges
}

//IsEdge return true if there is an edge between x and y
func (a AdjacencyListUndirectedGraph) IsEdge(x int, y int) bool {
	if x < 0 || y < 0 || y >= a.NbNodes || x >= a.NbNodes {
		return false
	}
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			for j := a.listNode[y]; j < a.listNode[y+1]; j++ {
				if a.succ[j] == x {
					return true
				}
			}
		}
	}
	return false
}

//RemoveEdge removes an edge (x,y) if exists
func (a *AdjacencyListUndirectedGraph) RemoveEdge(x int, y int) {
	if x < 0 || y < 0 || y >= a.NbNodes || x >= a.NbNodes || x == y {
		return
	}
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			a.NbEdges -= 1
			a.reduceNumberEdge(x, i)
			for j := a.listNode[y]; j < a.listNode[y+1]; j++ {
				if a.succ[j] == x {
					a.reduceNumberEdge(y, j)
					break
				}
			}
			break
		}
	}
}

//reduceNumberEdge delete an edge
func (a *AdjacencyListUndirectedGraph) reduceNumberEdge(nodePos, succPos int) {
	a.succ = append(a.succ[:succPos], a.succ[succPos+1:]...)
	a.weight = append(a.weight[:succPos], a.weight[succPos+1:]...)
	for nodePos = nodePos + 1; nodePos < len(a.listNode); nodePos++ {
		a.listNode[nodePos] -= 1
	}
}

//AddEdge add an edge (x,y) if not already present
func (a *AdjacencyListUndirectedGraph) AddEdge(x int, y int, p int) {
	if x < 0 || y < 0 || y >= a.NbNodes || x >= a.NbNodes || x == y {
		return
	}
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			return
		}
	}
	a.augmentNumberEdge(x, y, p)
	a.augmentNumberEdge(y, x, p)
	a.NbEdges += 1
}

//augmentNumberEdge add an edge
func (a *AdjacencyListUndirectedGraph) augmentNumberEdge(node int, succ int, p int) {
	ind := a.listNode[node]
	a.succ = append(a.succ, 0)
	copy(a.succ[ind+1:], a.succ[ind:])
	a.succ[ind] = succ

	a.weight = append(a.weight, 0)
	copy(a.weight[ind+1:], a.weight[ind:])
	a.weight[ind] = p

	for node = node + 1; node < len(a.listNode); node++ {
		a.listNode[node] += 1
	}
}

//GetNeighbors returns a new slice of int containing neighbors of node i
func (a AdjacencyListUndirectedGraph) GetNeighbors(node int) (neighbors []int) {
	for i := a.listNode[node]; i < a.listNode[node+1]; i++ {
		neighbors = append(neighbors, a.succ[i])
	}
	return neighbors
}

func (a AdjacencyListUndirectedGraph) GetWeight(x, y int) int {
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			return a.weight[i]
		}
	}
	return math.MaxInt64
}
