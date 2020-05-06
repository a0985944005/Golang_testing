package main

import (
	"testing"
)

func TestFizzBuzzSuccess(t *testing.T) {
	var testCases = []struct {
		in  []int
		out []string
	}{
		{
			in:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			out: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for idx := range testCases {
		for i := range testCases[idx].in {
			if testCases[idx].out[i] != FizzBuzz(testCases[idx].in[i]) {
				t.Errorf("not correct, input: %s ; output: %s", FizzBuzz(testCases[idx].in[i]), testCases[idx].out[i])
			}
		}
	}

}

func Benchmark_FizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(i)
	}
}
