package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func WalkWrapper(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	defer close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	
	go WalkWrapper(t1, ch1)
	go WalkWrapper(t2, ch2)
	
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break	
		}
	}
	
	return true
}

func main() {
	// Walk Test
	/*
	ch := make(chan int, 1)
	go WalkWrapper(tree.New(1), ch)
	for v := range ch {
		fmt.Printf("%v\n", v)
	}
	*/
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
