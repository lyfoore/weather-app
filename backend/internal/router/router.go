package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	weatherURL = "https://api.openweathermap.org/data/2.5/weather"
)

type Router struct{}

func StartRouter() {
	r := gin.Default()

	if err := godotenv.Load("../configs/.env"); err != nil {
		log.Fatal(err)
	}

	apiKey := os.Getenv("OWM_API_key")

	r.GET("/:cityName", func(ctx *gin.Context) {
		// apiKey :=
		// url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", weatherURL, ctx.Param("cityName"), apiKey)
		// http.Get(url)

		ctx.JSON(http.StatusOK, gin.H{
			ctx.Param("cityName"): "OK",
			"apiKey":              apiKey,
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
