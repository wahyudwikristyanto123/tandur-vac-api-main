package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tandur.com/src/adapter/out/mysql/user/adapter"
	"tandur.com/src/adapter/out/mysql/user/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/util"
)

func SetupUserController(r *gin.Engine) {
	repository := repository.NewUserRepository()
	adapter := adapter.NewUserAdapter(&repository)
	service := service.NewUserService(&adapter)

	r.GET("/user/token/:token", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		data, err := service.GetUserByToken(ctx.Params.ByName("token"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})
}
