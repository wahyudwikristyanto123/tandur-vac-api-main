package inbasket

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"tandur.com/src/adapter/out/mysql/inbasket/adapter"
	"tandur.com/src/adapter/out/mysql/inbasket/repository"
	"tandur.com/src/app/service"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

func SetupInbasketController(r *gin.Engine) {
	repository := repository.NewInbasketRepository()
	adapter := adapter.NewInbasketAdapter(&repository)
	service := service.NewInbasketService(&adapter)

	r.GET("/inbasket/token/:token", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var statuses []string = []string{}
		var q = ""
		if ctx.Query("status") == "" {
			status := "UNREAD,READ,DRAFT,SENT,REPLIED"
			statuses = strings.Split(status, ",")
		} else {
			status := ctx.Query("status")
			statuses = strings.Split(status, ",")
		}
		if ctx.Query("q") != "" {
			q = ctx.Query("q")
		}
		log.Println(statuses)
		data, err := service.GetByToken(q, ctx.Params.ByName("token"), statuses)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})

	r.GET("/inbasket/mailbox/:id", func(ctx *gin.Context) {
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

	r.GET("/inbasket/mailboxdraft/:id", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		id, err := strconv.ParseInt(ctx.Params.ByName("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		data, err := service.GetByParentIdAndDraft(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, data)
	})

	r.POST("/inbasket/reply", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var reqBody FullReplyRequest
		err := ctx.BindJSON(&reqBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err = service.FullReply(reqBody.ParentId, reqBody.Ccs, reqBody.Attachments, reqBody.Message)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})

	r.POST("/inbasket/replydraft", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var reqBody FullReplyRequest
		err := ctx.BindJSON(&reqBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		id, err := service.ReplyAndDraft(reqBody.ParentId, reqBody.Message)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	})

	r.POST("/inbasket/compose", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var reqBody domain.Mailbox
		err := ctx.BindJSON(&reqBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err = service.Create(reqBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})

	r.PATCH("/inbasket/updatedraft", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		var reqBody UpdateFullDraftRequest
		err := ctx.BindJSON(&reqBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err = service.UpdateAndDraft(reqBody.Id, reqBody.Message)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})

	r.PATCH("/inbasket/status", func(ctx *gin.Context) {
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
		err = service.UpdateStatus(reqBody.Id, reqBody.Status)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})

	r.DELETE("/inbasket/:id", func(ctx *gin.Context) {
		util.OpenMySQL()
		defer util.CloseMySQL()
		id, err := strconv.ParseInt(ctx.Params.ByName("token"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		err = service.DeleteMailbox(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})
}
