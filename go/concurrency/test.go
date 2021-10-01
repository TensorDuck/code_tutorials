package main

import (
	"fmt"
	"sync"
)

func get_n_random(i int, wg *sync.WaitGroup, answer *[][]int) {
	defer wg.Done()
	final := make([]int, i)
	for q := 0; q < i; q++ {
		final[q] = q
	}

	fmt.Println((final))
	(*answer)[i] = final
}

func main() {
	var wg sync.WaitGroup
	n_generate := 100
	store := make([][]int, n_generate)

	for i := 0; i < n_generate; i++ {
		wg.Add(1)
		go get_n_random(i, &wg, &store)
	}
	wg.Wait()
	fmt.Println(store)

}
