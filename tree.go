package main

import "fmt"

// Tree contains an identifier (value) and references to
// other trees (branches).
type Tree struct {
	value    int
	branches []Tree
}

func main() {
	root := Tree{}
	n := 4
	perm := PermSet(n)
	path := []int{}
	root.InitPermTree(perm)
	root.Prune(1)
	root.Traverse(path)
}

// PermSet returns the initial permutation {0, ..., n-1}.
func PermSet(permLen int) (perm []int) {
	perm = make([]int, permLen)
	for i := 0; i < permLen; i++ {
		perm[i] = i
	}
	return perm
}

// InitPermTree creates a fully grown tree consisting of all
// permutations on a given set.
func (t *Tree) InitPermTree(set []int) {
	cleanedSet := CleanSet(t.value, set)
	n := len(cleanedSet)
	if 0 < n {
		t.branches = make([]Tree, n)
		for i := 0; i < n; i++ {
			t.branches[i] = Tree{value: cleanedSet[i]}
			t.branches[i].InitPermTree(cleanedSet)
		}
	}
}

// CleanSet removes all values from a set and returns a new
// set.
func CleanSet(value int, set []int) (cleanedSet []int) {
	for i := 0; i < len(set); i++ {
		if set[i] != value {
			cleanedSet = append(cleanedSet, set[i])
		}
	}
	return cleanedSet
}

// Prune removes a tree by value from the branches of the
// current tree.
func (t *Tree) Prune(value int) {
	if index, contains := t.Contains(value); contains {
		t.branches = append(t.branches[:index], t.branches[index+1:]...)
	}
}

// Traverse TODO
func (t *Tree) Traverse(path []int) {
	if t.branches == nil {
		path = append(path, t.value)
		fmt.Println("PATH:", path)
	} else {
		for i := 0; i < len(t.branches); i++ {
			branchPath := append(path, t.value)
			t.branches[i].Traverse(branchPath)
		}
	}
}

// Contains returns the index and true if a tree exists in
// the current tree branches. If it doesn't exist, -1 and
// false is returned.
func (t *Tree) Contains(value int) (index int, contains bool) {
	index = -1
	for i, branch := range t.branches {
		if branch.value == value {
			index = i
			contains = true
			break
		}
	}
	return index, contains
}
