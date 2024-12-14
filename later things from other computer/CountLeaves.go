package main

import ( 
    "fmt"
    "strings"
    "strconv"
    "bufio"
    "os"
)

//Node is an object that represents a node of a tree.
type Node struct {
	Alignment      Alignment
	Num            int
	Age            float64
	Label          string
	Child1, Child2 *Node // if at leaf, both will be nil
}

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your CountLeaves() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
package main

import ( 
    "fmt"
    "strings"
    "strconv"
    "bufio"
    "os"
)

//Node is an object that represents a node of a tree.
type Node struct {
	Alignment      Alignment
	Num            int
	Age            float64
	Label          string
	Child1, Child2 *Node // if at leaf, both will be nil
}

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your CountLeaves() function here, along with any subroutines that you need. Please note the subroutines indicated in the problem description that are provided for you.
func CountLeaves(v *Node) int {
    if v.Child1 == nil && v.Child2 == nil {
        return 1
    } else if v.Child1 == nil {
        return CountLeaves(v.Child2)
    } else if v.Child2 == nil {
        return CountLeaves(v.Child1)
    }
    return CountLeaves(v.Child1) + CountLeaves(v.Child2)
}

