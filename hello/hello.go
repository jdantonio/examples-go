package main

import "fmt"

func main() {

	words := [...]string{"Hello, world", "你好，世界", "καλημ ́ρα κóσμ", "こんにちは世界"}

	fmt.Printf("%v or %v or %v or %v\n", words[0], words[1], words[2], words[3])
}
