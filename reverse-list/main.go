package main

import "fmt"

func main() {
	l := []int{1, 2, 3, 4}
	l = reverse[int](l)
	fmt.Println(l)
}

func reverse[T any](list []T) []T {

	for i, j := 0, len(list)-1; i < j; {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
	return list
}
