package inheritance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInheritance(t *testing.T) {
	animal := Animal{Mammal{"Bob", 10}}
	word := animal.Greet() // 等同于 animal.Mammal.Greet()，Mammal 为匿名字段，可省略不写。
	assert.Equal(t, "Hello, I'm Bob", word)
}
