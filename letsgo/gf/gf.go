// gf/gf.go
package main

import (
	"fmt"

	"github.com/karolgorecki/talks/letsgo/basement"
)

func main() {
	makeFries(basement.Potato)
}

func makeFries(potato string) {
	fmt.Println("GF makes french fries from delicous", potato)
}
