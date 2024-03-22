package main

import (
	"fmt"
	"testing"

	"github.com/nozzlium/heymat_backend/helper"
)

func TestMake(t *testing.T) {
	fmt.Println(byte(rune('0')))
}

func BenchmarkBench1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.GroupDecimals(50000000000)
	}
}

func BenchmarkBench2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helper.GroupDecimals2(50000000000)
	}
}
