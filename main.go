package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	"tandur.com/src/util"
)

var startTime = time.Now()

func main() {
	gin.DisableConsoleColor()

	// Ensure logs directory exists
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Printf("[Warning] Failed to create logs directory: %v", err)
	}

	logFile, err := os.OpenFile(fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	} else {
		gin.DefaultWriter = os.Stdout
	}

	// Initialize Database & Redis connection pools once at startup
	util.OpenMySQL()
	util.OpenRedis()

	r := gin.Default()
	r.Use(Cors())

	// Root & Health Check Endpoints
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Tandur VAC API is running",
			"status":  "healthy",
		})
	})

	r.GET("/health", func(ctx *gin.Context) {
		dbStatus := "connected"
		if db := util.GetMySQL(); db == nil || db.Ping() != nil {
			dbStatus = "disconnected"
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"status":    "UP",
			"database":  dbStatus,
			"uptime":    time.Since(startTime).String(),
			"timestamp": time.Now().Format(time.RFC3339),
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

	// Register Inbound REST Controllers
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

	// Port Resolution
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}
	serverAddr := fmt.Sprintf(":%s", port)

	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run HTTP Server in background goroutine for Graceful Shutdown
	go func() {
		log.Printf("[Server] Starting Tandur VAC API on port %s (http://localhost:%s)", port, port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[Server] Listen error: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("[Server] Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[Server] Server forced to shutdown: %v", err)
	}

	// Close database and cache pools
	util.CloseMySQL()
	util.CloseRedis()
	if logFile != nil {
		_ = logFile.Close()
	}

	log.Println("[Server] Server exited cleanly")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
