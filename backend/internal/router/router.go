package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyfoore/weather-app/configs"
)

const (
	weatherURL = "https://api.openweathermap.org/data/2.5/weather"
)

type Router struct{}

func StartRouter(cfg *configs.Config) {
	r := gin.Default()

	r.GET("/:cityName", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			ctx.Param("cityName"): "OK",
			"apiKey":              cfg.APIkey,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			ctx.Param("id"): "OK",
		})
	})

	r.GET("/api/time", func(ctx *gin.Context) {
		currentTime := time.Now().Format(time.RFC3339)
		ctx.JSON(http.StatusOK, gin.H{
			"time": currentTime,
		})
	})
	r.Run()
}
