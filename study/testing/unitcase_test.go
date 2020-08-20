package main

import "testing"

//单元测试
func TestAdd(t *testing.T) {
	var a = 10
	var b = 20

	c := Add(a, b)

	if c != 30 {
		t.Fatalf("invalid func Add()")
	}

	t.Logf("pass")
}

//单元测试
func TestSub(t *testing.T) {
	var a = 10
	var b = 20

	c := Sub(a, b)

	if c != -10 {
		t.Fatalf("invalid func Sub()")
	}

	t.Logf("pass")
}

//基准测试(压力测试)
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}
func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sub(10, 20)
	}
}
