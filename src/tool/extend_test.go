package tool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapAppend(t *testing.T) {
	from := map[int]int{
		1: 1, 2: 2,
	}
	to := map[int]int{
		3: 3, 4: 4,
	}
	to = MapAppend(from, to)
	assert.Equal(t, 1, to[1])
	assert.Equal(t, 2, to[2])
	assert.Equal(t, 3, to[3])
	assert.Equal(t, 4, to[4])
	to = nil
	to = MapAppend(from, to)
	assert.Equal(t, 1, to[1])
	assert.Equal(t, 2, to[2])
	assert.Equal(t, 0, to[3])
	assert.Equal(t, 0, to[4])

}
