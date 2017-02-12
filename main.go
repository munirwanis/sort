package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	runSorts(100)
	runSorts(1000)
	runSorts(10000)
	runSorts(100000)
	runSorts(1000000)
}

func runSorts(length int) {
	var array = rand.Perm(length)
	//fmt.Println(array)
	t0 := time.Now()
	bubbleSort(array)
	t1 := time.Now()
	//fmt.Println(array)
	fmt.Printf("%d length took %v\n", length, t1.Sub(t0))
}

func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func bubbleSort(arrayzor []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor)-1; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				swap(arrayzor, i, i+1)
				swapped = true
			}
		}
	}
}
