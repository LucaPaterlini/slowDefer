package main

import "testing"

func BenchmarkCallNoDefer(b *testing.B) {
	for i:=0;i<b.N;i++{
		CallNoDefer()
	}
}

func BenchmarkCallDefer(b *testing.B) {
	for i:=0;i<b.N;i++{
		CallDefer()
	}
}
