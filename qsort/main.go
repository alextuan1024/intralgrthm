package main

import "fmt"

func quicksort(array []int, p, r int) {
	if p < r {
		q := partition(array, p, r)
		quicksort(array, p, q-1)
		quicksort(array, q+1, r)
	}
}

func partition(arr []int, p, r int) int {
	var i, j, x int
	i = p - 1
	j = p
	x = arr[r]
	for ; j < r; j++ {
		if arr[j] <= x {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

func main() {
	tarr := []int{2, 8, 7, 1, 3, 5, 6, 4}
	quicksort(tarr, 0, 7)
	fmt.Println(tarr)
}
