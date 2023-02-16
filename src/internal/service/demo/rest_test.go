package demo

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var client = resty.New()

func TestRemoteMock(t *testing.T) {
	gock.New("http://test.com").
		Post("/first").
		Reply(http.StatusOK).BodyString("Hello")

	rtn, _ := client.R().Post("http://test.com/first")
	assert.Equal(t, "Hello", string(rtn.Body()))
}

func TestMain(m *testing.M) {
	// before test
	gock.InterceptClient(client.GetClient())
	defer gock.Off()
	defer gock.RestoreClient(client.GetClient())
	// test
	m.Run()
	// after test
}
