package main

import (
	"fmt"
)

func main() {
	w := map[int]string{
		0: "power",
		1: "go",
		2: "go",
		3: "ranges",
	}

	for _, v := range w {
		fmt.Println(v)
	}
}
