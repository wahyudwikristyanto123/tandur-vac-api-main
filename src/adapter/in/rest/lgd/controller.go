package problemanalysis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	assessorAdapter "tandur.com/src/adapter/out/mysql/assessor/adapter"
	assessorRepo "tandur.com/src/adapter/out/mysql/assessor/repository"
	"tandur.com/src/adapter/out/mysql/lgd/adapter"
	"tandur.com/src/adapter/out/mysql/lgd/repository"
	subtesAdapter "tandur.com/src/adapter/out/mysql/subtes/adapter"
	subtesRepo "tandur.com/src/adapter/out/mysql/subtes/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/util"
)

func SetupLgdController(r *gin.Engine) {
	repository := repository.NewLgdRepository()
	assessorRepository := assessorRepo.NewAssessorRepository()
	subtesRepository := subtesRepo.NewSubtesRepository()

	adapter := adapter.NewLgdAdapter(&repository)
	assessorAdapter := assessorAdapter.NewAssessorAdapter(&assessorRepository)
	subtesAdapter := subtesAdapter.NewSubtesAdapter(&subtesRepository)

	service := service.NewLgdService(&adapter, &assessorAdapter, &subtesAdapter)

	r.GET("/lgd/token/:token", func(ctx *gin.Context) {
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
}
