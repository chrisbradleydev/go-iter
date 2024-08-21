package main

import (
	"fmt"
)

func main() {
	s := []string{"the", "quick", "brown", "fox", "jumped", "over", "the", "fence"}
	for i, v := range Backwards(s) {
		fmt.Println(i, v)
	}
}

func Backwards(s []string) func(func(int, string) bool) {
	return func(yeild func(int, string) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yeild(i, s[i]) {
				return
			}
		}
	}
}
