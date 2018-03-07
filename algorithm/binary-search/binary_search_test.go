package binarysearch

import "testing"
import "github.com/stretchr/testify/assert"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 4, 6, 7}
	index := BinarySearch(nums, 2)
	assert.Equal(t, index, 1)
}

func BenchmarkBinarySearch(b *testing.B) {
	nums := []int{1, 2, 4, 6, 7}
	for i := 0; i < b.N; i++ {
		BinarySearch(nums, 2)
	}
}
