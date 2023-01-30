package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnum(t *testing.T) {
	mySeasons := []Season{Spring, Summer, Autumn, Winter, 100}
	assert.Equal(t, true, mySeasons[0].IsValid())
	assert.Equal(t, true, mySeasons[1].IsValid())
	assert.Equal(t, true, mySeasons[2].IsValid())
	assert.Equal(t, true, mySeasons[3].IsValid())
	assert.Equal(t, false, mySeasons[4].IsValid())
}
