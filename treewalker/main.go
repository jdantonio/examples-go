package main

// Exercise: Equivalent Binary Trees
// https://tour.golang.org/concurrency/8

import (
	"fmt"
	"golang.org/x/tour/tree"
)

//type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var a1, a2 []int
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		Walk(t1, ch1)
	}()

	go func() {
		Walk(t2, ch2)
	}()

	for i := 0; i < 10; i++ {
		a1 = append(a1, <-ch1)
	}
	close(ch1)

	for i := 0; i < 10; i++ {
		a2 = append(a2, <-ch2)
	}
	close(ch2)

	return compare(a1, a2)
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {
	if Same(tree.New(1), tree.New(1)) == true {
		fmt.Println("Success!")
	} else {
		fmt.Println("Fail!")
	}

	if Same(tree.New(1), tree.New(2)) == false {
		fmt.Println("Success!")
	} else {
		fmt.Println("Fail!")
	}
}
