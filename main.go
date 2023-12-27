package main

import (
	"fmt"
	"math/rand"
)

func main() {

	names := []string{"you", "me", "someone", "he", "she"}

	min := 0
	max := len(names) - 1

	fmt.Println(names[rand.Intn(max-min+1)+min])

}
