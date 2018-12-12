package balancedparentheses

import "testing"
import "github.com/stretchr/testify/assert"

func TestBinarySearch(t *testing.T) {
	assert.True(t, IsParanthesisBalanced("[()]{}{[()()]()}"))
	assert.False(t, IsParanthesisBalanced("[(])"))
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsParanthesisBalanced("[()]{}{[()()]()}")
	}
}
