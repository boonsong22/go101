package main

func main() {
	println(squareArea(10))
	println(power(2, 2))
}

func squareArea(a int) int {
	return a * a
}

func prime(n int) {
	for i := 0; i < n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			println(count)
		}
	}
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
