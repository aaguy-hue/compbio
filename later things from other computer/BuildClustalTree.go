package main

type Tree []*Node

// Alignment is a multiple alignment object corresponding to a slice of strings
type Alignment []string

// Node is an object that represents a node of a tree.
// We also think of a node as a "cluster" when building a UPGMA tree.
type Node struct {
	Alignment      Alignment
	Num            int
	Age            float64
	Label          string
	Child1, Child2 *Node // if at leaf, both will be nil
}

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your BuildClustalTree() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func BuildClustalTree(guideTree Tree, leafStrings []string, match float64, mismatch float64, gap float64, supergap float64) Tree {
	InitializeGuideTree(guideTree, leafStrings)
}

// initializeGuideTree takes a Tree object and a collection of strings as input.
// It sets the alignment of each leaf equal to a blank alignment
// consisting of the string labeling that leaf.
func InitializeGuideTree(guide Tree, strs []string) {
	for idx, str := range strs {
		guide[idx].Alignment = Alignment{str}
	}
}
