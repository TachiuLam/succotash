package main

import "fmt"

func main() {
	var a, b int
	b = add(a)
	fmt.Println(b)
}

func add(a int) (b int) {
	defer func() {
		a++
		b++
	}()
	a++
	return a
}
