package graph

func ExplorGraphUndirectedInWidth(graph IUndirectedGraph, startingNode int) (nodes []int) {
	mark := make([]bool, graph.GetNbNodes())
	var toVisit []int
	mark[startingNode] = true
	toVisit = append(toVisit, startingNode)
	nodes = append(nodes, startingNode)
	for len(toVisit) > 0 {
		nodeInspected := toVisit[0]
		toVisit = append(toVisit[:0], toVisit[1:]...)
		for _, v := range graph.GetNeighbors(nodeInspected) {
			if mark[v] == false {
				mark[v] = true
				nodes = append(nodes, v)
				toVisit = append(toVisit, v)
			}
		}
	}
	return nodes
}

func ExplorGraphDirectedInWidth(graph IDirectedGraph, startingNode int) (nodes []int) {
	mark := make([]bool, graph.GetNbNodes())
	var toVisit []int
	mark[startingNode] = true
	toVisit = append(toVisit, startingNode)
	nodes = append(nodes, startingNode)
	for len(toVisit) > 0 {
		nodeInspected := toVisit[0]
		toVisit = append(toVisit[:0], toVisit[1:]...)
		for _, v := range graph.GetSuccessors(nodeInspected) {
			if mark[v] == false {
				mark[v] = true
				nodes = append(nodes, v)
				toVisit = append(toVisit, v)
			}
		}
	}
	return nodes
}

func ExplorerGraphDirectedInDepth(graph IDirectedGraph, sommet int, visited []int) []int {
	visited = append(visited, sommet)
	for _, succ := range graph.GetSuccessors(sommet) {
		isIn := false
		for i := 0; i < len(visited); i++ {
			if visited[i] == succ {
				isIn = true
			}
		}
		if !isIn {
			visited = ExplorerGraphDirectedInDepth(graph, succ, visited)
		}
	}
	return visited
}

func ExplorerGraphUndirectedInDepth(graph IUndirectedGraph, sommet int, visited []int) []int {
	visited = append(visited, sommet)
	for _, succ := range graph.GetNeighbors(sommet) {
		isIn := false
		for i := 0; i < len(visited); i++ {
			if visited[i] == succ {
				isIn = true
			}
		}
		if !isIn {
			visited = ExplorerGraphUndirectedInDepth(graph, succ, visited)
		}
	}
	return visited
}

func FirstPathChecker(graph IUndirectedGraph, sommet int) (paths [][]int) {
	var atteint []int
	atteint = append(atteint, sommet)
	paths = make([][]int, graph.GetNbNodes())
	for i := 0; i < graph.GetNbNodes(); i++ {
		paths[i] = make([]int, 0)
	}
	for _, s := range graph.GetNeighbors(sommet) {
		isAtteint := false
		for _, val := range atteint {
			if val == s {
				isAtteint = true
				break
			}
		}
		if !isAtteint {
			explorerSommet(graph, s, sommet, atteint, paths)
		}
	}
	return paths
}

func explorerSommet(graph IUndirectedGraph, visite int, pred int, atteint []int, paths [][]int) {
	atteint = append(atteint, visite)
	paths[visite] = append(paths[visite], paths[pred]...)
	paths[visite] = append(paths[visite], pred)
	for _, t := range graph.GetNeighbors(visite) {
		isAtteint := false
		for _, val := range atteint {
			if val == t {
				isAtteint = true
				break
			}
		}
		if !isAtteint {
			explorerSommet(graph, t, visite, atteint, paths)
		}
	}
}
