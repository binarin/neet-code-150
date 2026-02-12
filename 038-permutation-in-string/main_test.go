package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	assert.Equal(t, true, checkInclusion("ab", "eidbaooo"))
	assert.Equal(t, false, checkInclusion("ab", "eidboaoo"))
}
