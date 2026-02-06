package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, "6", multiply("2", "3"))
	assert.Equal(t, "0", multiply("0", "0"))
	assert.Equal(t, "56088", multiply("123", "456"))
}

func TestAddSlices(t *testing.T) {
	// 1 + 2 = 3
	assert.Equal(t, []byte{3}, addSlices([]byte{1}, []byte{2}))

	// 5 + 7 = 12 (with carry)
	assert.Equal(t, []byte{2, 1}, addSlices([]byte{5}, []byte{7}))

	// 9 + 91 = 100 (different lengths)
	assert.Equal(t, []byte{0, 0, 1}, addSlices([]byte{9}, []byte{1, 9}))

	// 99 + 1 = 100 (carry propagation)
	assert.Equal(t, []byte{0, 0, 1}, addSlices([]byte{9, 9}, []byte{1}))

	// 123 + 456 = 579
	assert.Equal(t, []byte{9, 7, 5}, addSlices([]byte{3, 2, 1}, []byte{6, 5, 4}))

	// 999 + 999 = 1998
	assert.Equal(t, []byte{8, 9, 9, 1}, addSlices([]byte{9, 9, 9}, []byte{9, 9, 9}))

	// 0 + 0 = 0
	assert.Equal(t, []byte{0}, addSlices([]byte{0}, []byte{0}))

	// 0 + 123 = 123
	assert.Equal(t, []byte{3, 2, 1}, addSlices([]byte{0}, []byte{3, 2, 1}))

	// empty slices
	assert.Equal(t, []byte{}, addSlices([]byte{}, []byte{}))
}

func TestMul1(t *testing.T) {
	// 2 * 3 = 6
	assert.Equal(t, []byte{6}, mul1([]byte{2}, 3))

	// 5 * 3 = 15 (with carry)
	assert.Equal(t, []byte{5, 1}, mul1([]byte{5}, 3))

	// 12 * 3 = 36
	assert.Equal(t, []byte{6, 3}, mul1([]byte{2, 1}, 3))

	// 99 * 9 = 891
	assert.Equal(t, []byte{1, 9, 8}, mul1([]byte{9, 9}, 9))

	// 123 * 4 = 492
	assert.Equal(t, []byte{2, 9, 4}, mul1([]byte{3, 2, 1}, 4))

	// 999 * 9 = 8991
	assert.Equal(t, []byte{1, 9, 9, 8}, mul1([]byte{9, 9, 9}, 9))

	// multiply by 0 (preserves length, no trimming)
	assert.Equal(t, []byte{0, 0, 0}, mul1([]byte{5, 4, 3}, 0))

	// multiply by 1
	assert.Equal(t, []byte{3, 2, 1}, mul1([]byte{3, 2, 1}, 1))

	// empty slice
	assert.Equal(t, []byte{}, mul1([]byte{}, 5))
}
