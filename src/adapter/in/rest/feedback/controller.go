package feedback

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tandur.com/src/adapter/out/mysql/feedback/adapter"
	"tandur.com/src/adapter/out/mysql/feedback/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/util"
)

func SetupFeedbackController(r *gin.Engine) {
	repository := repository.NewFeedbackRepository()
	adapter := adapter.NewFeedbackAdapter(&repository)
	service := service.NewFeedbackService(&adapter)

	r.POST("/feedback/submit", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var reqBody FeedbackRequest
		err := ctx.BindJSON(&reqBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err = service.SubmitFeedback(RequestToDomain(reqBody))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})
}
