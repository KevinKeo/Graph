package graph

type AdjacencyListUndirectedGraph struct {
	listNode, succ []int
}

func NewAdjacencyListUndirectedGraph(generatedGraph [][]int) *AdjacencyListUndirectedGraph {
	nodes := make([]int, len(generatedGraph)+1)
	var succ []int
	nodes[0] = 0
	for i, v := range generatedGraph {
		for j, v2 := range v {
			if v2 == 1 {
				succ = append(succ, j)
			}
			nodes[i+1] = len(succ)
		}
	}
	return &AdjacencyListUndirectedGraph{nodes, succ}
}

func (a AdjacencyListUndirectedGraph) GetNbNodes() int {
	return len(a.listNode) - 1
}

func (a AdjacencyListUndirectedGraph) ToAdjacencyMatrix() [][]int {
	matrix := make([][]int, a.GetNbNodes())
	for i := 0; i < a.GetNbNodes(); i++ {
		matrix[i] = make([]int, a.GetNbNodes())
		for j := a.listNode[i]; j < a.listNode[i+1]; j++ {
			matrix[i][a.succ[j]] = 1
		}
	}
	return matrix
}

func (a AdjacencyListUndirectedGraph) GetNbEdges() int {
	return len(a.succ) / 2
}

func (a AdjacencyListUndirectedGraph) IsEdge(x int, y int) bool {
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			return true
		}
	}
	return false
}

func (a *AdjacencyListUndirectedGraph) RemoveEdge(x int, y int) {
	if x == y {
		return
	}
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
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

func (a *AdjacencyListUndirectedGraph) reduceNumberEdge(nodePos, succPos int) {
	a.succ = append(a.succ[:succPos], a.succ[succPos+1:]...)
	for nodePos = nodePos + 1; nodePos < len(a.listNode); nodePos++ {
		a.listNode[nodePos] -= 1
	}
}

func (a *AdjacencyListUndirectedGraph) AddEdge(x int, y int) {
	if x == y {
		return
	}
	for i := a.listNode[x]; i < a.listNode[x+1]; i++ {
		if a.succ[i] == y {
			return
		}
	}
	a.augmentNumberEdge(x)
	a.augmentNumberEdge(y)

}

func (a *AdjacencyListUndirectedGraph) augmentNumberEdge(x int) {
	ind := a.listNode[x]
	a.succ = append(a.succ, 0)
	copy(a.succ[ind+1:], a.succ[ind:])
	a.succ[ind] = 1
	for x = x + 1; x < len(a.listNode); x++ {
		a.listNode[x] += 1
	}
}

/**

    GetNbNodes() int
    ToAdjacencyMatrix() [][]int
    GetNbEdges() int
    IsEdge(int, int) bool
    RemoveEdge(int, int)
    AddEdge(int, int)
    GetNeighbors(int) []int
**/
