package exercises

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

func (t *Tree) Height() (h int) {
	if t.leftBranch != nil || t.rightBranch != nil{
		h +=
	}
	return
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
