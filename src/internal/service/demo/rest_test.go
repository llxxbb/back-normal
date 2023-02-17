package demo

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"gitlab.cdel.local/platform/go/platform-common/old"
)

var client = resty.New()

func TestServerTime(t *testing.T) {
	// 注意 gock 拦截时，resty 的 BaseUrl 不会附加到请求路径上，
	// 所以下面的 New 必须用空串。
	gock.New("").
		Post("/cdel@+/server/time").
		Reply(http.StatusOK).
		// BodyString("Hello")
		JSON(map[string]any{
			"success": true,
			"result":  1676540186616,
		})

	raw, err := getTime(client)
	assert.Nil(t, err)
	var result old.ServiceResult
	json.Unmarshal(raw.Body(), &result)
	assert.Equal(t, float64(1676540186616), result.Result)
	assert.Equal(t, true, result.Success)
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
