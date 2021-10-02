package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var recursiveStep func(t *tree.Tree)
	recursiveStep = func(t *tree.Tree) {
	   if t == nil {
		  return
	   }
	   recursiveStep(t.Left)
	   ch <- t.Value
	   recursiveStep(t.Right)
	}
	recursiveStep(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2, ok2 := <- ch2
		if v1 != v2 || !ok2 {
			return false
		}
	}
	_, ok2 := <- ch2
	return !ok2
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
