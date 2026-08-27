package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tandur.com/src/adapter/out/mysql/user/adapter"
	"tandur.com/src/adapter/out/mysql/user/repository"
	"tandur.com/src/app/service"
)

func SetupUserController(r *gin.Engine) {
	repository := repository.NewUserRepository()
	adapter := adapter.NewUserAdapter(&repository)
	service := service.NewUserService(&adapter)

	r.GET("/user/token/:token", func(ctx *gin.Context) {
		token := ctx.Params.ByName("token")
		data, err := service.GetUserByToken(token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})
}
