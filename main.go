package main

import (
	"fmt"

	"github.com/hsmtkk/improved-happiness/leap"
)

func main() {
	years := []int{1900, 1996, 1997, 2000}
	for _, year := range years {
		if leap.IsLeap(year) {
			fmt.Printf("%d is leap year\n", year)
		} else {
			fmt.Printf("%d is NOT leap year\n", year)
		}
	}
}
