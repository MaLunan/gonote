package main

import (
	"fmt"
)

func jiujiu() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Printf("%d*%d=%d\n", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%d", j, i, i*j)
			}

		}
	}
}
func main() {
	jiujiu()
}
