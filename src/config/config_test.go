package config

import (
	"os"
	"reflect"
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

func TestFindField(t *testing.T) {
	c := Config{}
	c.Mysql.DBName = "lxb"
	root := reflect.ValueOf(&c).Elem()
	to := findField(&root, "Mysql.DBName")
	assert.Equal(t, "lxb", to.String())
}
