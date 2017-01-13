package main

import (
	"fmt"
	. "github.com/kevin/Graph/graphutils"
)

func main() {

	graph := GenerateGraphData(4, 4, true)
	c := NewAdjacencyMatrixUndirectedGraph(graph)
	fmt.Println("MatrixGraph")
	printGraph(c.ToAdjacencyMatrix())
	adjacencyList := NewAdjacencyListUndirectedGraphWithMatrix(c.ToAdjacencyMatrix())
	fmt.Println(adjacencyList)
	fmt.Println("graph adjancy list")
	printGraph(adjacencyList.ToAdjacencyMatrix())
	fmt.Println(c.IsEdge(0, 2))
	fmt.Println("remove before", adjacencyList.NbEdges)
	adjacencyList.RemoveEdge(0, 1)
	fmt.Println("remove after", adjacencyList.NbEdges)
	fmt.Println(adjacencyList)
	printGraph(adjacencyList.ToAdjacencyMatrix())
	fmt.Println("//////////////////////////////////////")
	fmt.Println("add")
	fmt.Println(adjacencyList)
	fmt.Println("add before", adjacencyList.NbEdges)
	adjacencyList.AddEdge(0, 1)
	fmt.Println("add after", adjacencyList.NbEdges)
	fmt.Println(adjacencyList)
	printGraph(adjacencyList.ToAdjacencyMatrix())

	var in IUndirectedGraph
	in = c
	fmt.Println("Neighbors :", in.GetNeighbors(0))
	in = adjacencyList
	fmt.Println("Neighbors :", in.GetNeighbors(0))

	/*fmt.Println("/////////////////////////////////")
	in = c
	printGraph(c.Matrice)
	fmt.Println("test")
	c.AddEdge(0, 2)
	printGraph(c.Matrice)
	fmt.Println("Neighbors :", in.GetNeighbors(0))
	fmt.Println("test")
	c.RemoveEdge(0, 2)
	printGraph(c.Matrice)*/
	graph = GenerateGraphData(4, 6, false)
	c = NewAdjacencyMatrixUndirectedGraph(graph)
	fmt.Println(graph)
	listDirected := NewAdjacencyListDirectedGraphWithMatrix(graph)
	fmt.Println("matrix")
	printGraph(c.ToAdjacencyMatrix())
	fmt.Println("list")
	printGraph(listDirected.ToAdjacencyMatrix())
	fmt.Println(listDirected.GetSuccessors(1))
	fmt.Println(listDirected.GetPredecessors(3))
	fmt.Println("listinverse")
	printGraph(listDirected.ComputeInverse().ToAdjacencyMatrix())

	fmt.Println("list")
	listDirected.RemoveArc(0, 1)
	printGraph(listDirected.ToAdjacencyMatrix())

	newList := NewAdjacencyListDirectedGraphWithInterface(listDirected)

}

func printGraph(g [][]int) {
	for _, v := range g {
		fmt.Println(v)
	}

}
