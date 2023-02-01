package main

import (
	"net/http"

	controllers "github.com/fransedor/golang-spotify/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var access_token string = ""
var refresh_token string = ""

func main() {
	godotenv.Load()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	router.GET("/login", controllers.Login)
	router.GET("/callback", controllers.Callback(&access_token, &refresh_token))
	router.GET("/tracks", controllers.GetTracks(&access_token, &refresh_token))
	router.Run("localhost:8080")
}
