package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gama",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, v := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, v), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("second thing to print", &wg)
}
