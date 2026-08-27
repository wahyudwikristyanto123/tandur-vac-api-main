package problemanalysis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	assessorAdapter "tandur.com/src/adapter/out/mysql/assessor/adapter"
	assessorRepo "tandur.com/src/adapter/out/mysql/assessor/repository"
	"tandur.com/src/adapter/out/mysql/one_on_one/adapter"
	"tandur.com/src/adapter/out/mysql/one_on_one/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/util"
)

func SetupOneOnOneController(r *gin.Engine) {
	repository := repository.NewOneOnOneRepository()
	assessorRepository := assessorRepo.NewAssessorRepository()
	adapter := adapter.NewOneOnOneAdapter(&repository)
	assessorAdapter := assessorAdapter.NewAssessorAdapter(&assessorRepository)
	service := service.NewOneOnOneService(&adapter, &assessorAdapter)

	r.GET("/1on1/token/:token", func(ctx *gin.Context) {
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
