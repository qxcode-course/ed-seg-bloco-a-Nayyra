package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if b == 0 {
		return a
	}

	if a == 0 {
		return b
	}

	r := a % b
	return mdc(b,r)

}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
