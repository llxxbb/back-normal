package tool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type JsonTester struct {
	A string
}

func TestObject2String(t *testing.T) {
	rtn, err := Object2String(nil)
	assert.Equal(t, "", rtn)
	assert.Nil(t, err)

	rtn, err = Object2String(123)
	assert.Equal(t, "123", rtn)
	assert.Nil(t, err)

	rtn, err = Object2String(JsonTester{A: "hello"})
	assert.Equal(t, "{\"A\":\"hello\"}", rtn)
	assert.Nil(t, err)
}

func TestString2Object(t *testing.T) {
	object, err := String2Object[int]("", true)
	assert.Nil(t, object)
	assert.Nil(t, err)

	object, err = String2Object[int]("123", true)
	assert.Equal(t, 123, *object)
	assert.Nil(t, err)

	jt, err := String2Object[JsonTester]("{\"A\":\"hello\"}", true)
	assert.Equal(t, JsonTester{A: "hello"}, *jt)
	assert.Nil(t, err)
}
