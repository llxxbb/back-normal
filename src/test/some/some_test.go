package some

import (
	"back/demo/test"
	"testing"

	"github.com/stretchr/testify/assert"
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
