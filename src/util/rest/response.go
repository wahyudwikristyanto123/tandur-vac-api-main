package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReturnSuccess(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

func ReturnError(ctx *gin.Context, code int, err error) {
	if err != nil {
		log.Printf("%+v", err)
		ctx.JSON(code, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
}

func GetParam(ctx *gin.Context, name string) string {
	return ctx.Params.ByName(name)
}

func GetQueryParamInt64(ctx *gin.Context, name string) (int64, error) {
	return strconv.ParseInt(ctx.Query(name), 0, 64)
}

func GetQueryParamString(ctx *gin.Context, name string) string {
	return ctx.Query(name)
}
