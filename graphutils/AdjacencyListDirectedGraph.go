/*
Package graph provides a generic framework to manipulate simple graph
Containing representation of different simple type of graphs and implements method to work with them
*/
package graph

//AdjacencyListDirectGraph represents an directed graph in the form of an adjacency List
type AdjacencyListDirectedGraph struct {
	NbNodes, NbArcs int
	listNode, succ  []int
}

//NewAdjacencyListDirectGraph create a new AdjacencyListDirectGraph's pointer
func NewAdjacencyListDirectGraph(generatedGraph [][]int) *AdjacencyListDirectedGraph {
	nodes := make([]int, len(generatedGraph)+1)
	var succ []int
	for i, array := range generatedGraph {
		for j, v2 := range array {
			if v2 == 1 {
				succ = append(succ, j)
			}
		}
		nodes[i+1] = len(succ)
	}
	return &AdjacencyListDirectedGraph{len(nodes) - 1, len(succ), nodes, succ}
}

func (a AdjacencyListDirectedGraph) ToAdjacencyMatrix() [][]int {
	matrix := make([][]int, a.GetNbNodes())

	for i := 0; i < a.GetNbNodes(); i++ {
		matrix[i] = make([]int, a.NbNodes)
		for j := a.listNode[i]; j < a.listNode[i+1]; j++ {
			matrix[i][a.succ[j]] = 1
		}
	}
	return matrix
}

//GetNbNodes return the number of Nodes of the graph
func (a AdjacencyListDirectedGraph) GetNbNodes() int {
	return a.NbNodes
}

//GetNbArcs gives the number of arcs in the graph
func (a AdjacencyListDirectedGraph) GetNbArcs() int {
	return a.NbArcs
}

//IsArc return true if there is an arc from x to y
func (a AdjacencyListDirectedGraph) IsArc(x int, y int) bool {
	if x < 0 || y < 0 || y > a.NbNodes || x > a.NbNodes {
		return false
	}

	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			return true
		}
	}
	return false
}

//RemoveArc removes an arc from x to y if exists
func (a *AdjacencyListDirectedGraph) RemoveArc(x int, y int) {
	if x < 0 || y < 0 || y > a.NbNodes || x > a.NbNodes || x == y {
		return
	}

	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			a.NbArcs -= 1
			a.reduceNumberArc(x, i)
			break
		}
	}
}

//reduceNumberArc delete an edge
func (a *AdjacencyListDirectedGraph) reduceNumberArc(nodePos, succPos int) {
	a.succ = append(a.succ[:succPos], a.succ[succPos+1:]...)
	for nodePos = nodePos + 1; nodePos < len(a.listNode); nodePos++ {
		a.listNode[nodePos] -= 1
	}
}

//AddArc add an arc from x to y if not already present
func (a *AdjacencyListDirectedGraph) AddArc(x int, y int) {
	if x < 0 || y < 0 || y > a.NbNodes || x > a.NbNodes || x == y {
		return
	}
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			return
		}
	}
	a.augmentNumberArc(x, y)
	a.NbArcs += 1
}

//augmentNumberArc add an arc from x to y
func (a *AdjacencyListDirectedGraph) augmentNumberArc(node int, succ int) {
	ind := a.listNode[node]
	a.succ = append(a.succ, 0)
	copy(a.succ[ind+1:], a.succ[ind:])
	a.succ[ind] = succ

	for node = node + 1; node < len(a.listNode); node++ {
		a.listNode[node] += 1
	}
}

//GetSuccessors returns a new int representing successors of node x
func (a AdjacencyListDirectedGraph) GetSuccessors(x int) (successors []int) {
	if x < 0 || x > a.NbNodes {
		return successors
	}
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		successors = append(successors, a.succ[i])
	}
	return successors
}

//GetPredecessors returns a new int representing predecessors of node x
func (a AdjacencyListDirectedGraph) GetPredecessors(x int) (pred []int) {
	if x < 0 || x > a.NbNodes {
		return pred
	}
	var position []int
	for i, v := range a.succ {
		if v == x {
			position = append(position, i)
		}
	}

	for _, valPos := range position {
		for i := 1; i < len(a.listNode); i++ {
			if valPos < a.listNode[i] {
				pred = append(pred, i-1)
				break
			}
		}
	}
	return pred
}

//ComputeInverse returns the inverse of the graph
func (a AdjacencyListDirectedGraph) ComputeInverse() IDirectedGraph {
	var succ []int
	nodes := make([]int, a.NbNodes+1)
	for i := 0; i < a.NbNodes; i++ {
		succ = append(succ, a.GetPredecessors(i)...)
		nodes[i+1] = len(succ)
	}
	return &AdjacencyListDirectedGraph{len(nodes) - 1, len(succ), nodes, succ}

}
