package subtes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"tandur.com/src/adapter/out/mysql/subtes/adapter"
	"tandur.com/src/adapter/out/mysql/subtes/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/util"
)

func SetupSubtesController(r *gin.Engine) {
	repository := repository.NewSubtesRepository()
	adapter := adapter.NewSubtesAdapter(&repository)
	service := service.NewSubtesService(&adapter)

	r.GET("/subtes/token/:token", func(ctx *gin.Context) {
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

	r.GET("/subtes/result/:id", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		data, err := service.GetResultById(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})

	r.PATCH("/subtes/status", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var reqBody UpdateRequest
		err := ctx.BindJSON(&reqBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err = service.UpdateStatusById(reqBody.Id, reqBody.Status)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})

	r.POST("/subtes/submit", func(ctx *gin.Context) {
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
		err = service.SubmitResult(reqBody.Id, reqBody.Result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})
}
