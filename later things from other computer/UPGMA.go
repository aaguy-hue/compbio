package main

import (
	"strconv"
)

// DistanceMatrix is a 2D slice of floats
type DistanceMatrix [][]float64

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

// Insert your UPGMA() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree {
	numLeaves := len(mtx)
	tree := InitializeTree(speciesNames)
	clusters := InitializeClusters(tree)
	for p := numLeaves; p < 2*numLeaves-1; p++ {
		row, col, val := FindMinElement(mtx)
		tree[p].Age = val / 2
		tree[p].Child1 = clusters[row]
		tree[p].Child2 = clusters[col]
		clusterSize1 := CountLeaves(tree[p].Child1)
		clusterSize2 := CountLeaves(tree[p].Child2)
		mtx = AddRowCol(row, col, clusterSize1, clusterSize2, mtx)
		mtx = DeleteRowCol(mtx, row, col)
		clusters = append(clusters, tree[p])
		clusters = DeleteClusters(clusters, row, col)
	}
	return tree
}

// helper functions
func InitializeTree(speciesNames []string) Tree {
	numLeaves := len(speciesNames)
	//var t Tree // a Tree is []*Node
	var t Tree
	t = make([]*Node, 2*numLeaves-1)
	// all of these pointers have default value of nil; we need to point them at nodes

	// we should create 2n-1 nodes.
	for i := range t {
		var vx Node
		// let's label the first numLeaves nodes with the appropriate species name.
		// by default, vx.age = 0.0, and its children are nil.
		if i < numLeaves {
			//at a leaf ... let's assign its label.
			vx.Label = speciesNames[i]
		} else {
			// let's just give it an unspecific name
			vx.Label = "Ancestor species " + strconv.Itoa(i)
		}
		// indexing the node
		vx.Num = i
		// one thing to do: point t[i] at vx
		t[i] = &vx
	}

	return t
}

func InitializeClusters(t Tree) []*Node {
	numNodes := len(t)
	numLeaves := (numNodes + 1) / 2
	clusters := make([]*Node, numLeaves)
	// clusters[i] should point to the i-th leaf node of t
	for i := range clusters {
		clusters[i] = t[i]
	}

	return clusters
}

func DeleteClusters(clusters []*Node, row, col int) []*Node {
	clusters = append(clusters[:col], clusters[col+1:]...)
	clusters = append(clusters[:row], clusters[row+1:]...)
	return clusters
}
