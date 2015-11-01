package main

import "fmt"

type Void interface{}

type Predicate func(Void) bool

func Select(collection []Void, f Predicate) []Void {
	var selected []Void
	for _, item := range collection {
		if f(item) {
			selected = append(selected, item)
		}
	}
	return selected
}

func main() {
	c := []Void{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := Select(c, func(i Void) bool { x, _ := i.(int); return x%2 == 0 })
	fmt.Println(s)
}
