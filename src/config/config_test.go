package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_default(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, "back-normal", c.ProjectName)
	assert.Equal(t, "v0.0.1", c.ProjectVersion)
	assert.Equal(t, "product", c.Env)
	wd, _ := os.Getwd()
	assert.Equal(t, wd, c.WorkPath)
	assert.Equal(t, true, len(c.Host) > 0)
	println(c.LogPath)
	assert.Equal(t, true, len(c.LogPath) > 0)
}

func TestFindField(t *testing.T) {
	c := Config{}
	c.Mysql.DBName = "lxb"
	root := reflect.ValueOf(&c).Elem()
	to := findField(&root, "Mysql.DBName")
	assert.Equal(t, "lxb", to.String())
}
