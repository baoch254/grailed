package ginauth

import (
	"net/http"

	"grailed/libs/gr-api/data-access-user/usermodel"
	"grailed/libs/gr-api/data-access-user/userstorage"
	"grailed/libs/gr-api/feature-auth/authbiz"
	common "grailed/libs/gr-api/shared-common"
	hasher "grailed/libs/gr-api/utils-hasher"

	"github.com/gin-gonic/gin"
)

func Register(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := authbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
