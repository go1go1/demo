package main

import "testing"

func BenchmarkBubbleSort(b *testing.B) {
	var a = [8]int{8, 3, 2, 9, 4, 6, 10, 0}
	for i := 0; i < b.N; i++ {
		bubbleSort(a)
	}
}
