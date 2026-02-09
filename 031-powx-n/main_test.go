package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyPow(t *testing.T) {
	assert.InDelta(t, 1024.00000, myPow(2.00000, 10), 1e-5)
	assert.InDelta(t, 9.26100, myPow(2.10000, 3), 1e-5)
	assert.InDelta(t, 0.25000, myPow(2.00000, -2), 1e-5)
}
