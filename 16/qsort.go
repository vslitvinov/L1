package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// quicksort - sorts array with O(n*log(n))
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	rand.Seed(time.Now().UnixNano())
	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func main() {
	a := []int{3, 7, 2, 9, 0, 1, 6, 8, 4}
	fmt.Println(a)
	fmt.Println(quicksort(a))
	sort.Ints(a)
	fmt.Println(a)
}
