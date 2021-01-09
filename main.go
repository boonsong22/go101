package main

import (
	"go101/fizzbuzz"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/fizzbuzz/:number", func(ctx *gin.Context) {
		number := ctx.Param("number")
		n, err := strconv.Atoi(number)
		if (err) != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"say": fizzbuzz.Say(n),
		})
	})

	router.Run()

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

func couple(s string) []string {
	result := []string{}
	s += "*"
	for ; len(s) > 1; s = s[2:] {

	}
	return result
}
