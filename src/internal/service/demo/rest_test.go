package demo

import (
	"cdel/demo/Normal/internal/dao"
	"cdel/demo/Normal/internal/entity"
	"cdel/demo/Normal/tool"
	"github.com/goccy/go-json"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"gitlab.cdel.local/platform/go/platform-common/access"
	"gitlab.cdel.local/platform/go/platform-common/old"
)

// used for TestDbSelect --------------------------
func TestDbSelect(t *testing.T) {
	ctrl := gomock.NewController(t)
	// 创建 mock 对象, 下面的 NewMockTmpTableDaoI 由 mockgen 创建
	// 在 src 目录下运行 mockgen.exe -source=".\internal\dao\tmp.go" -destination=".\internal\dao\tmp_mock.go" -package=dao
	m := dao.NewMockTmpTableDaoI(ctrl)
	m.EXPECT().SelectByName(gomock.Any(), gomock.Eq("tom")).Return(
		[]entity.TmpTable{}, nil,
	).Times(1)
	tmpDao = m

	// mock gin.Context
	req := access.ParaIn[string]{Data: "tom"}
	tool.GinCall(req, DbSelect)
}

// will be intercepted
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
	var result old.ServiceResult[int64]
	_ = json.Unmarshal(raw.Body(), &result)
	assert.Equal(t, int64(1676540186616), result.Result)
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
