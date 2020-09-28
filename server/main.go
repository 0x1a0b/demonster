package main

import (
	_ "github.com/0x1a0b/demonster/server/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
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
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	MountApi(r.Group("/api"))
	r.NoRoute(NotFound)
	return
}
// @title Demonster API
// @version 1.0
// @description This is an example api to demonstrate things
// @termsOfService http://localhost:8080/tos.html

// @contact.name Demonster API Support
// @contact.url https://github.com/0x1a0b/demonster/issues
// @contact.email support@localhost.local

// @license.name MIT
// @license.url https://github.com/0x1a0b/demonster/blob/master/LICENSE

// @host localhost:8080
// @BasePath /api
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

type ActionsResponse struct {
	BodyLen int    `json:"body_len"`
	Message string `json:"message"`
	Body    string `json:"body"`
	Headers interface{} `json:"headers"`
}

func ActionsRequest(c *gin.Context) {
	rawData, _ := c.GetRawData()
	body := string(rawData[:])
	length := len(rawData)
	headers := c.Request.Header

	ar := ActionsResponse{
		BodyLen: length,
		Message: "success",
		Body: body,
		Headers: headers,
	}
	c.JSON(http.StatusOK, ar)
	return
}
