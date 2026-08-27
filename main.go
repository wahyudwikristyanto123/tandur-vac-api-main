package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cv "tandur.com/src/adapter/in/rest/cv"
	feedback "tandur.com/src/adapter/in/rest/feedback"
	inbasket "tandur.com/src/adapter/in/rest/inbasket"
	lgd "tandur.com/src/adapter/in/rest/lgd"
	ononone "tandur.com/src/adapter/in/rest/one_on_one"
	problemanalysis "tandur.com/src/adapter/in/rest/problem_analysis"
	qna "tandur.com/src/adapter/in/rest/qna"
	roleplay "tandur.com/src/adapter/in/rest/roleplay"
	"tandur.com/src/adapter/in/rest/subtes"
	upload "tandur.com/src/adapter/in/rest/upload"
	"tandur.com/src/adapter/in/rest/user"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create(fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02")))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	r.Use(Cors())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello World",
		})
	})

	r.GET("/benchmark", func(ctx *gin.Context) {
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename=benchmark.txt")
		ctx.Header("Content-Type", "application/txt")
		ctx.Header("Content-Length", "0")
		ctx.File("./benchmark.txt")
	})

	subtes.SetupSubtesController(r)
	problemanalysis.SetupProblemAnalysisController(r)
	lgd.SetupLgdController(r)
	qna.SetupQnaController(r)
	inbasket.SetupInbasketController(r)
	user.SetupUserController(r)
	ononone.SetupOneOnOneController(r)
	upload.SetupUploadController(r)
	roleplay.SetupRoleplayController(r)
	cv.SetupCvController(r)
	feedback.SetupFeedbackController(r)
	r.Run()
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
