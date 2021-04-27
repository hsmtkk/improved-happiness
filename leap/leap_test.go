package leap_test

import (
	"testing"

	"github.com/hsmtkk/improved-happiness/leap"
	"github.com/stretchr/testify/assert"
)

func TestIsLeap(t *testing.T) {
	assert.False(t, leap.IsLeap(1997))
	assert.True(t, leap.IsLeap(1996))
	assert.False(t, leap.IsLeap(1900))
	assert.True(t, leap.IsLeap(2000))
}
