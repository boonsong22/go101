package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n Int = 10
	fmt.Printf("%q\n", n.String())
	n.Set(29)
	fmt.Printf("%q\n", n.String())
}

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int) {
	*i = Int(n)
}
