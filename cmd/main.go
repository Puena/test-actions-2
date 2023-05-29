package main

import (
	"fmt"

	"github.com/Puena/test-actions-2/internal/calculation"
)

func main() {
	fmt.Println("hello world!")
	sum := calculation.Add(1, 2)
	fmt.Println("calculation of 1 and 2: ", sum)
}
