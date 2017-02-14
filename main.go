package main

import (
	"fmt"
	. "github.com/kevin/Graph/graphutils"
	"math"
)

func main() {

	matD := [][]int{{math.MaxInt64, 1, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, 1, 1, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}}
	matU := [][]int{{math.MaxInt64, 1, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1}, {1, math.MaxInt64, math.MaxInt64, 1, 1, math.MaxInt64, math.MaxInt64}, {1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, 1, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {math.MaxInt64, math.MaxInt64, 1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}, {1, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}}
	directedGraph := NewAdjacencyListDirectedGraphWithMatrix(matD)
	undirectedGraph := NewAdjacencyListUndirectedGraphWithMatrix(matU)
	fmt.Println(ExplorGraphDirectedInWidth(directedGraph, 0), "test")
	fmt.Println(ExplorGraphUndirectedInWidth(undirectedGraph, 0), "test")
	fmt.Println(ExplorerGraphDirectedInDepth(directedGraph, 0, make([]int, 0)))
	fmt.Println(ExplorerGraphUndirectedInDepth(undirectedGraph, 0, make([]int, 0)))

	binarytree := NewTree()
	binarytree.AddElem(6)
	binarytree.AddElem(8)
	binarytree.AddElem(4)
	binarytree.AddElem(5)
	binarytree.PrintTree()
	binarytree.DeleteFirstElem()

	binarytree.PrintTree()

	binarytree.AddElem(1)
	binarytree.AddElem(2)
	binarytree.AddElem(3)
	binarytree.AddElem(1)
	binarytree.AddElem(2)
	binarytree.AddElem(3)
	binarytree.PrintTree()

	binarytree.DeleteFirstElem()

	binarytree.DeleteFirstElem()

	binarytree.DeleteFirstElem()

	binarytree.DeleteFirstElem()

	binarytree.PrintTree()
	binarytree.DeleteFirstElem()
	binarytree.PrintTree()
	fmt.Println(FirstPathChecker(undirectedGraph, 0))
}

func printGraph(g [][]int) {
	for _, v := range g {
		fmt.Println(v)
	}

}
