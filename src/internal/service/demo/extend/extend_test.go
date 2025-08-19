package extend

import (
	"back/demo/internal/service/demo/inheritance"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrement(t *testing.T) {
	var i = 10
	// 调用扩展的方法
	rtn := MyInt(i).increment()
	assert.Equal(t, 11, rtn)
}

func TestHello(t *testing.T) {
	dog := MyAnimal{inheritance.Mammal{Name: "dog", Age: 3}}
	// 调用扩展的方法
	rtn := dog.hello()
	assert.Equal(t, "dog: hello", rtn)
}
