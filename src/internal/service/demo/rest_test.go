package demo

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestRemoteMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

}
