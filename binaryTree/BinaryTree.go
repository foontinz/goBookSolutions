package main

import (
	"fmt"
	"strings"
)

// Tree left - small, right - big
type Tree struct {
	rightBranch *Tree
	leftBranch  *Tree
	Key         int
}

func (t *Tree) Insert(k int) {

	if k < t.Key {
		if t.leftBranch == nil {
			t.leftBranch = &Tree{Key: k}
		} else {
			t.leftBranch.Insert(k)
		}
	} else if k > t.Key {
		if t.rightBranch == nil {
			t.rightBranch = &Tree{Key: k}
		} else {
			t.rightBranch.Insert(k)
		}
	}
}
func (t *Tree) Search(k int) bool {
	if k == t.Key {
		return true
	}
	if k < t.Key {
		if t.leftBranch != nil {
			return t.leftBranch.Search(k)
		}
	} else if k > t.Key {
		if t.rightBranch != nil {
			return t.rightBranch.Search(k)
		}

	}

	return false
}

func (t *Tree) InOrderTraversal() []int {
	resArr := []int{}

	if t.rightBranch != nil {
		resArr = append(resArr, t.rightBranch.InOrderTraversal()...)
	}
	resArr = append(resArr, t.Key)
	if t.leftBranch != nil {
		resArr = append(resArr, t.leftBranch.InOrderTraversal()...)
	}
	return resArr

}
func (t *Tree) Height() (h int) {
	if t == nil {
		return 0
	}
	if t.leftBranch == nil && t.rightBranch == nil {
		return 1
	}

	lHeight := t.leftBranch.Height()
	rHeight := t.rightBranch.Height()

	if lHeight >= rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}

func PrintTree(t *Tree) {
	fmt.Printf("%v%s", t.Key, strings.Repeat("\t", t.Height()))
	if t.leftBranch != nil {
		fmt.Print("")
		PrintTree(t.leftBranch)
	}
	if t.rightBranch != nil {
		PrintTree(t.rightBranch)
	}
}
