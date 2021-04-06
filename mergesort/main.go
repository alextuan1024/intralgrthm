package mergesort

const sentinel = 0x7FF0000000000000

func merge(arr []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	la := make([]int, n1+1)
	ra := make([]int, n2+1)

	var i, j int
	for i = 0; i < n1; i++ {
		la[i] = arr[p+i]
	}
	for j = 0; j < n2; j++ {
		ra[j] = arr[q+j+1]
	}
	la[n1] = sentinel
	ra[n2] = sentinel
	i = 0
	j = 0
	for k := p; k <= r; k++ {
		if la[i] <= ra[j] {
			arr[k] = la[i]
			i += 1
		} else {
			arr[k] = ra[j]
			j += 1
		}
	}
}

func mergesort(array []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergesort(array, p, q)
		mergesort(array, q+1, r)
		merge(array, p, q, r)
	}
}
