package main

import (
	"fmt"
	"iter"
)

func Backward(s []string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func main() {
	s := []string{"hello", "world"}
	for i, x := range Backward(s) {
		fmt.Println(i, x)
	}
}
