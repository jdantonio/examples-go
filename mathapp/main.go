package main

import (
	"fmt"
	"github.com/jdantonio/examples.go/hello"
	"github.com/jdantonio/examples.go/mymath"
)

func main() {
	hello.World()
	fmt.Printf("Sqrt(2) = %v\n", mymath.Sqrt(2))
}
