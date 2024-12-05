package some

import (
	"cdel/demo/Normal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthCheck(t *testing.T) {
	assert.True(t, true)
}

func TestMain(t *testing.M) {
	test.InitTestEnv("..")
	// some other init begin
	// ...
	// some other init emd
	t.Run()
}
