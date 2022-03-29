package ginauth

import (
	"net/http"

	"grailed/libs/gr-api/data-access-user/usermodel"
	"grailed/libs/gr-api/data-access-user/userstorage"
	"grailed/libs/gr-api/feature-auth/authbiz"
	common "grailed/libs/gr-api/shared-common"
	hasher "grailed/libs/gr-api/utils-hasher"
	"grailed/libs/gr-api/utils-token-provider/jwt"

	"github.com/gin-gonic/gin"
)

func Login(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data usermodel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		biz := authbiz.NewLoginBiz(store, tokenProvider, md5, 60*60*24*30)

		account, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
