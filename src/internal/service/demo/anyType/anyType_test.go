package anyType

import (
	"back/demo/internal/service/demo/inheritance"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyType(t *testing.T) {
	rtn := PrintType("hello")
	assert.Equal(t, "type: string, value: hello", rtn)
	rtn = PrintType(inheritance.Mammal{Name: "Katty", Age: 2})
	assert.Equal(t, "name: Mammal, kind: struct, string: inheritance.Mammal", rtn)
}
