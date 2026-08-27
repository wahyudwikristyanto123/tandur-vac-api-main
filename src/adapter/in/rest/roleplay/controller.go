package roleplay

import (
	"net/http"

	"github.com/gin-gonic/gin"
	assessorAdapter "tandur.com/src/adapter/out/mysql/assessor/adapter"
	assessorRepo "tandur.com/src/adapter/out/mysql/assessor/repository"
	"tandur.com/src/adapter/out/mysql/roleplay/adapter"
	"tandur.com/src/adapter/out/mysql/roleplay/repository"
	subtesAdapter "tandur.com/src/adapter/out/mysql/subtes/adapter"
	subtesRepo "tandur.com/src/adapter/out/mysql/subtes/repository"
	"tandur.com/src/app/service"
)

func SetupRoleplayController(r *gin.Engine) {
	repository := repository.NewRoleplayRepository()
	assessorRepository := assessorRepo.NewAssessorRepository()
	subtesRepository := subtesRepo.NewSubtesRepository()

	adapter := adapter.NewRoleplayAdapter(&repository)
	assessorAdapter := assessorAdapter.NewAssessorAdapter(&assessorRepository)
	subtesAdapter := subtesAdapter.NewSubtesAdapter(&subtesRepository)

	service := service.NewRoleplayService(&adapter, &assessorAdapter, &subtesAdapter)

	r.GET("/roleplay/token/:token", func(ctx *gin.Context) {
		token := ctx.Params.ByName("token")
		data, err := service.GetByToken(token)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})
}
