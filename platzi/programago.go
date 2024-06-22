package main

import "fmt"

func main() {
	var j int = 1
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n", j)
		j = j + 1
	}
}
