package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var traceId string

func GetLogTraceId() string {
	traceId = uuid.NewString()
	return traceId
}

func INFO(message string) {
	logger := gin.DefaultWriter
	now := time.Now().Format("2006-01-02 15:04:05")
	logger.Write([]byte(fmt.Sprintf("[INFO]\t%s\t%s\t%s\n", now, traceId, message)))
}

func WARN(message string) {
	logger := gin.DefaultWriter
	now := time.Now().Format("2006-01-02 15:04:05")
	logger.Write([]byte(fmt.Sprintf("[WARN]\t%s\t%s\t%s\n", now, traceId, message)))
}

func ERROR(message string) {
	logger := gin.DefaultErrorWriter
	now := time.Now().Format("2006-01-02 15:04:05")
	logger.Write([]byte(fmt.Sprintf("[ERROR]\t%s\t%s\t%s\n", now, traceId, message)))
}

func DEBUG(message string) {
	logger := gin.DefaultWriter
	now := time.Now().Format("2006-01-02 15:04:05")
	logger.Write([]byte(fmt.Sprintf("[DEBUG]\t%s\t%s\t%s\n", now, traceId, message)))
}
