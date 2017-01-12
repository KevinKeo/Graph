/*
Package graph which provides a generic framework to manipulate simple graph
*/
package graph

import (
	"math/rand"
	"time"
)

/*
IGraph Interface to manipulate graph
*/
type IGraph interface {
	GetNbNodes() int
	ToAdjacencyMatrix() [][]int
}

/*
IUndirectedGraph Interface to manipulate non oriented graph
*/
type IUndirectedGraph interface {
	IGraph
	GetNbEdges() int
	IsEdge(int, int) bool
	RemoveEdge(int, int)
	AddEdge(int, int)
	GetNeighbors(int) []int
}

//IDirectedGraph interface to manipulate oriented graph
type IDirectedGraph interface {
	IGraph
	GetNbArcs() int
	IsArc(int, int) bool
	RemoveArc(int, int)
	AddArc(int, int)
	GetSuccessors(int) []int
	GetPredecessors(int) []int
	ComputeInverse() IDirectedGraph
}

/*
GenerateGraphData generate a graph of n nodes, with m connection between the node, with an option s to be symetric and render it by a matrice of int
*/
func GenerateGraphData(n int, m int, s bool) (graph [][]int) {
	type couple struct {
		i, j int
	}
	graph = make([][]int, n)
	var list []couple
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if s && j > i {
				list = append(list, couple{i, j})
			} else if !s && j != i {
				list = append(list, couple{i, j})
			}
		}
	}

	for ; m > 0; m-- {
		s1 := rand.NewSource(time.Now().UnixNano() * int64(m+1))
		r1 := rand.New(s1)
		i := r1.Intn(len(list))
		graph[list[i].i][list[i].j] = 1
		if s {
			graph[list[i].j][list[i].i] = 1
		}
		list = append(list[:i], list[i+1:]...)
	}

	return graph

}

//GenerateGraphDataBis is another implementation of generate graph
func GenerateGraphDataBis(n int, m int, s bool) (graph [][]int) {
	graph = make([][]int, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	attempt := n * n
	if s == true {
		attempt = attempt / 2
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			attempt -= 1
			if i == j {
				graph[i][j] = 0
			} else {
				r := 0
				if m > 0 {
					s1 := rand.NewSource((time.Now().UnixNano() - int64(m)) * (time.Now().UnixNano() + int64(m)))
					r1 := rand.New(s1)
					r = r1.Intn(2)
					if attempt <= m {
						r = 1
					}
					m -= r
				}
				if s == true && i > j {
					m += r
					graph[i][j] = graph[j][i]
				} else {
					graph[i][j] = r
				}

			}
		}
	}
	return graph
}

//toSymetric transform an matrice into an symetric matrice
func toSymetric(matrice [][]int) (newmatrice [][]int) {
	newmatrice = make([][]int, len(matrice))
	for i := range newmatrice {
		newmatrice[i] = make([]int, len(matrice[i]))
	}

	for i, v1 := range matrice {
		for j := range v1 {
			newmatrice[i][j] = matrice[i][j]
			newmatrice[j][i] = matrice[i][j]
		}
	}
	return newmatrice
}
