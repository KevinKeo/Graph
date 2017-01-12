package graph

type AdjacencyMatrixUndirectedGraph struct {
	Matrice [][]int
}

/**
func NewAdjacencyMatrixUndirectedGraph(IUndirectedGraph) {

	}**/
func (a AdjacencyMatrixUndirectedGraph) GetNbEdges() (n int) {
	matrice := a.Matrice
	for i, v1 := range matrice {
		for j, v2 := range v1 {
			if j > i {
				n += v2
			}
		}
	}
	return n
}

func (a AdjacencyMatrixUndirectedGraph) GetNbNodes() (n int) {
	n = len(a.Matrice)
	return n
}

func (a AdjacencyMatrixUndirectedGraph) ToAdjacencyMatrix() [][]int {
	return a.Matrice
}

func (a AdjacencyMatrixUndirectedGraph) IsEdge(i int, j int) bool {
	return a.Matrice[i][j] == 1
}

func (a *AdjacencyMatrixUndirectedGraph) AddEdge(i int, j int) {
	if i >= len(a.Matrice) || j >= len(a.Matrice) {
		return
	}
	a.Matrice[i][j] = 1
	a.Matrice[j][i] = 1
}

func (a *AdjacencyMatrixUndirectedGraph) RemoveEdge(i int, j int) {
	if i >= len(a.Matrice) || j >= len(a.Matrice) {
		return
	}
	a.Matrice[i][j] = 0
	a.Matrice[j][i] = 0
	/*
	   a.Matrice = GenerateGraphData(1, 0, false)
	   a.Matrice[i][j] = 7
	   a.Matrice[j][i] = 7*/

}
func (a AdjacencyMatrixUndirectedGraph) GetNeighbors(i int) []int {
	return a.Matrice[i]
}
