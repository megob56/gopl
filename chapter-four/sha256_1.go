package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	count := 0
	for i := range c1 {
		for j := range c2 {
			if i != j {
				count++
			}
		}
	}

	fmt.Println(count)
}
