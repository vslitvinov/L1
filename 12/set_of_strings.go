package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

// set - remove all duplicate from slice
func set(s []string) []string {
	var res []string
	setOfString := make(map[string]bool, len(s))

	for _, val := range s {
		setOfString[val] = true
	}

	for key := range setOfString {
		res = append(res, key)
	}

	return res
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	setArr := set(arr)

	fmt.Printf("Arr: %v\nSet: %v\n", arr, setArr)
}
