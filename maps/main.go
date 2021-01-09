package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./data.csv")
	r := csv.NewReader(f)
	records, _ := r.ReadAll()

	for _, v := range records {
		fmt.Println(v)

	}
}
