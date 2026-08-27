package problemanalysis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tandur.com/src/adapter/out/mysql/problem_analysis/adapter"
	"tandur.com/src/adapter/out/mysql/problem_analysis/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/util"
)

func SetupProblemAnalysisController(r *gin.Engine) {
	repository := repository.NewProblemAnalysisRepository()
	adapter := adapter.NewProblemAnalysisAdapter(&repository)
	service := service.NewProblemAnalysisService(&adapter)

	r.GET("/pa/token/:token", func(ctx *gin.Context) {
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
