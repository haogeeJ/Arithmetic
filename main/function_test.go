package main

import "testing"

func BenchmarkGenerate(b *testing.B) {
	n = 10000
	r = 10

	for i := 0; i < b.N; i++ {
		generateProblems()
	}

}

func BenchmarkChekc(b *testing.B) {
	exercisefile = "exercisefile.txt"
	answerfile = "answerfile.txt"
	for i := 0; i < b.N; i++ {
		check()
	}
}
