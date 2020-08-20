package main

import "testing"

func BenchmarkGo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go1()
	}
}

func BenchmarkGo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go2()
	}
}
