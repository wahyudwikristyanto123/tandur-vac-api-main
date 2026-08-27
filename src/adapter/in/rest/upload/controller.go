package upload

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"tandur.com/src/util/logger"
)

func SetupUploadController(r *gin.Engine) {

	r.POST("/upload", func(ctx *gin.Context) {
		file, handler, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		defer file.Close()
		logger.INFO(fmt.Sprintf("Uploaded File: %+v\n", handler.Filename))
		logger.INFO(fmt.Sprintf("File Size: %+v\n", handler.Size))
		logger.INFO(fmt.Sprintf("MIME Header: %+v\n", handler.Header))

		buf := bytes.NewBuffer(nil)
		_, err = io.Copy(buf, file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		url := fmt.Sprintf("https://sg.storage.bunnycdn.com/tandur-dev/asesi/%s", handler.Filename)
		req, err := http.NewRequest(http.MethodPut, url, buf)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		req.Header.Set("AccessKey", os.Getenv("CDN_API_KEY"))
		req.Header.Set("Content-Type", "application/octet-stream")
		req.Header.Set("accept", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		bodyString := string(bodyBytes)

		if resp.StatusCode == 401 {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": bodyString,
			})
			return
		}

		if resp.StatusCode == 400 {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": bodyString,
			})
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"url": fmt.Sprintf("https://tandur.b-cdn.net/asesi/%s", handler.Filename),
		})
	})
}
