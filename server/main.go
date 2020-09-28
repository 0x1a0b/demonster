package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := Setup()
	r.Run("0.0.0.0:8080")
}

func Setup() (r *gin.Engine) {
	r = gin.Default()
	r.Use(CORS)
	MountApi(r.Group("/api"))
	r.NoRoute(NotFound)
	return
}

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
	return
}

func MountApi(router *gin.RouterGroup) {
	router.GET("/alive", AlivePing)
	actions := router.Group("/actions")
	{
		actions.POST("/request", ActionsRequest)
	}
	return
}

func AlivePing(c *gin.Context) {
	now := time.Now()
	c.JSON(http.StatusOK, gin.H{"status": "ok", "time": now.String()})
	return
}

func NotFound(c *gin.Context) {
	path := c.Request.URL.Path
	method := c.Request.Method
	client := c.ClientIP()
	log.WithFields(log.Fields{
		"path":    path,
		"methods": method,
		"client":  client,
	}).Debug("path not found handler")
	c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	return
}

func ActionsRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "status": "success" })
	return
}