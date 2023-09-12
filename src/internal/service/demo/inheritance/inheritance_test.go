package inheritance

import (
	json2 "github.com/goccy/go-json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInheritance(t *testing.T) {
	animal := Animal{Mammal{"Bob", 10}}
	// 等同于 animal.Mammal.Greet()，Mammal 为匿名字段，可省略不写。
	word := animal.Greet()
	assert.Equal(t, "Hello, I'm Bob", word)
}

func TestJson(t *testing.T) {
	animal := Animal{Mammal{"Bob", 10}}
	// 等同于 animal.Mammal.Greet()，Mammal 为匿名字段，可省略不写。
	marshal, _ := json2.Marshal(animal)
	json := string(marshal)
	assert.Equal(t, "{\"Name\":\"Bob\",\"Age\":10}", json)
}
