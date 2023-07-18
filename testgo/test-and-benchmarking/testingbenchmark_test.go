package main

import (
	"fmt"
	"testing"
)

func Intmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntminbasic(t *testing.T) {
	ans := Intmin(2, -2)
	if ans != -2 {
		t.Errorf("intmin(2,-2) = %d, want -2", ans)
	}
}

func TestIntmintabledriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -1, -1},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d , %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			if ans := Intmin(tt.a, tt.b); ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntmin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Intmin(1, 2)
	}
}
