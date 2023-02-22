package main

import (
	"net/http"

	controllers "github.com/fransedor/golang-spotify/controllers"
	"github.com/fransedor/golang-spotify/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var access_token string = ""
var refresh_token string = ""

func main() {
	godotenv.Load()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
	}))
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	router.GET("/login", controllers.Login)
	router.GET("/callback", controllers.Callback(&access_token, &refresh_token))
	router.GET("/tracks", controllers.GetTracks(&access_token, &refresh_token))
	router.GET("/profile", controllers.GetProfile)
	router.GET("/devices", controllers.GetAvailableDevices)
	router.PUT("/playback/start", middlewares.GetActiveDevice, controllers.StartOrPausePlayback)
	router.GET("/playback/current", middlewares.GetActiveDevice, controllers.GetCurrentlyPlaying)
	router.Run("localhost:8080")
}
