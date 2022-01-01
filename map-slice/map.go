package main

import "fmt"

func main() {
	scores := map[string]int {
		"ming": 10,
		"zhang": 12,
	}

	//_, _ := scores["tang"]

	for k, v := range scores {
		fmt.Println(k, " - ", v)
	}

	delete(scores, "zhang")

	multiSlice := make([][]int, 2)
	multiSlice = append(multiSlice, []int{1, 2, 3})
	multiSlice = append(multiSlice, []int{4, 5, 6})
}