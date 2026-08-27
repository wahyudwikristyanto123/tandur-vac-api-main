package cv

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tandur.com/src/adapter/out/mysql/cv/adapter"
	"tandur.com/src/adapter/out/mysql/cv/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/util"
)

func SetupCvController(r *gin.Engine) {
	repository := repository.NewCvRepository()
	adapter := adapter.NewCvAdapter(&repository)
	service := service.NewCvService(&adapter)

	r.GET("/cv/token/:token", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		data, err := service.GetByToken(ctx.Params.ByName("token"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})

	r.POST("/cv/submit", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var reqBody SubmitRequest
		err := ctx.BindJSON(&reqBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err = service.SubmitCv(RequestToDomain(reqBody))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})
}
