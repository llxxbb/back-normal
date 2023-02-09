package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_default(t *testing.T) {
	assert.Equal(t, "back-normal", C.ProjectName)
	assert.Equal(t, "v0.0.1", C.ProjectVersion)
	assert.Equal(t, "product", C.Env)
	wd, _ := os.Getwd()
	assert.Equal(t, wd, C.WorkPath)
	assert.Equal(t, true, len(C.Host) > 0)
	println(C.LogPath)
	assert.Equal(t, true, len(C.LogPath) > 0)
}
