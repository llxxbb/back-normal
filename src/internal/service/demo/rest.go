package demo

import (
	"cdel/demo/Normal/config"
	"net/http"

	"gitlab.cdel.local/platform/go/platform-common/access"
	"gitlab.cdel.local/platform/go/platform-common/def"
	"gitlab.cdel.local/platform/go/platform-common/old"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func V1(c *gin.Context) {
	in := old.Request{}
	c.ShouldBindJSON(&in)
	out := old.GetSuccess(in.Params)
	c.JSON(http.StatusOK, out)
}
func V2(c *gin.Context) {
	in := access.ParaIn{}
	c.ShouldBindJSON(&in)
	out := access.GetSuccessResult(in.Data)
	c.JSON(http.StatusOK, out)
}

func DbSelect(c *gin.Context) {
	in := access.ParaIn{}
	c.ShouldBindJSON(&in)
	rows, err := config.CTX.DemoDB.Query("SELECT * FROM album WHERE artist = ? LIMIt ?", in.Data)
	if err != nil {
		zap.S().Warn(err)
		c.JSON(http.StatusOK, access.GetErrorResultD(def.ET_ENV, def.E_ENV.Code, def.E_ENV.Msg+err.Error(), nil))
	}
	defer rows.Close()
	// for rows.Next() {
	// 	var alb Album
	// 	if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
	// 		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	// 	}
	// 	albums = append(albums, alb)
	// }
}
