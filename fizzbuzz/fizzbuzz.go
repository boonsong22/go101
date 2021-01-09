package fizzbuzz

import (
	"math/rand"
	"strconv"
)

func Say(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)

}

func String() string {
	r1 := rand.Intn(100)
	r2 := rand.Intn(100)
	r3 := rand.Intn(100)
	r4 := rand.Intn(100)
	return Say(r1) + Say(r2) + Say(r3) + Say(r4)
}
