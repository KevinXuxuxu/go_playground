package main

func partition(a []int, i, j int) int {
	p, l, r := i, i+1, j-1
	for l < r {
		for l < len(a) && a[l] <= a[p] {
			l++
		}
		if l >= r || l >= len(a) {
			break
		}
		for r > 0 && a[r] > a[p] {
			r--
		}
		if l >= r || r <= 0 {
			break
		}
		a[r], a[l] = a[l], a[r]
		r--
		l++
	}
	for l >= len(a) || a[l] > a[p] {
		l--
	}
	a[l], a[p] = a[p], a[l]
	return l
}

func _quick_sort(a []int, i, j int) {
	if j-i <= 1 {
		return
	}
	p := partition(a, i, j)
	_quick_sort(a, i, p)
	_quick_sort(a, p+1, j)
}

func quick_sort(a []int) {
	_quick_sort(a, 0, len(a))
}
