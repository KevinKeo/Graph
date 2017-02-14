package graph

import "fmt"

type Tree struct {
	valeur []int
}

func NewTree() *Tree {
	var valeur []int
	return &Tree{valeur}
}

func (t *Tree) AddElem(n int) {
	t.valeur = append(t.valeur, n)
	i := len(t.valeur) - 1
	for ; i > 0 && t.valeur[(i-1)/2] > n; i = (i - 1) / 2 {
		t.valeur[i] = t.valeur[(i-1)/2]
		t.valeur[(i-1)/2] = n
	}
}

func (t *Tree) DeleteFirstElem() {
	length := len(t.valeur) - 1
	if length < 0 {
		return
	}
	t.valeur[0] = t.valeur[length]
	t.valeur = t.valeur[:length]
	r := 1
	l := 2

	for i := 0; r < length; {
		leaf := r
		if l < length {
			if t.valeur[leaf] > t.valeur[l] {
				leaf = l
			}
		}
		if t.valeur[leaf] <= t.valeur[i] {
			tmp := t.valeur[i]
			t.valeur[i] = t.valeur[leaf]
			t.valeur[leaf] = tmp
			i = leaf
		} else {
			break
		}

		r = 2*i + 1
		l = r + 1

	}
}

/*
func (graph IUndirectedGraph) AlgoPrim() {
	var en []couple

}*/
func (t *Tree) PrintTree() {
	fmt.Println(t.valeur)
}
