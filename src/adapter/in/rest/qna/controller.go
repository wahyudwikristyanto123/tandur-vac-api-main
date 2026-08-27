package qna

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"tandur.com/src/adapter/out/mysql/qna/adapter"
	"tandur.com/src/adapter/out/mysql/qna/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

func SetupQnaController(r *gin.Engine) {
	repository := repository.NewQnaRepository()
	adapter := adapter.NewQnaAdapter(&repository)
	service := service.NewQnaService(&adapter)

	r.GET("/qna/token/:token", func(ctx *gin.Context) {
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

	r.GET("/qna/id/:id", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		data, err := service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})

	r.POST("/qna/result", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var data domain.QnaResultRequest
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err := service.UpsertResult(data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
		})
	})

	r.POST("/qna/results", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var data []domain.QnaResultRequest
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err := service.UpsertResults(data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
		})
	})

	r.GET("/qna/results/token/:token", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		data, err := service.GetResultsByToken(ctx.Params.ByName("token"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})
}
