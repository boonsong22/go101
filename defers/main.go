package main

import "fmt"

func main() {
	deferPrint(2)
}

func deferPrint(n int) {
	defer fmt.Println(n)
	defer func() { fmt.Println(n) }()
	n *= n
	fmt.Println(n)
}
