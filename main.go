package main

import (
	"fmt"
	. "github.com/kevin/graphe/graphutils"
)

func main() {
	//var in IUndirectedGraph
	graph := GenerateGraphData(4, 4, true)
	c := AdjacencyMatrixUndirectedGraph{graph}
	fmt.Println(c.IsEdge(0, 2))
	adjacencyList := NewAdjacencyListUndirectedGraph(graph)
	fmt.Println(adjacencyList)
	fmt.Println("graph adjancy list")
	printGraph(adjacencyList.ToAdjacencyMatrix())
	fmt.Println("remove")
	adjacencyList.RemoveEdge(0, 2)
	fmt.Println(adjacencyList)
	printGraph(adjacencyList.ToAdjacencyMatrix())
	fmt.Println("//////////////////////////////////////")
	fmt.Println("add")
	fmt.Println(adjacencyList)
	adjacencyList.AddEdge(0, 2)
	printGraph(adjacencyList.ToAdjacencyMatrix())
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

}

func printGraph(g [][]int) {
	for _, v := range g {
		fmt.Println(v)
	}

}
