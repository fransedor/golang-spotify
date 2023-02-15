package controllers

import (
	"log"
	"os"

	helper "github.com/fransedor/golang-spotify/helpers"
	"github.com/gin-gonic/gin"
)

func StartPlayback(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	spotifyURL := os.Getenv("SPOTIFY_URL")
	req, err := helper.CreateHTTPRequestWithHeader("PUT", spotifyURL+"/me/player/play", nil, authHeader)
	if err != nil {
		log.Fatal(err)
	}
	var response string
	var status string
	status, err = helper.GetHTTPResponse(req, &response)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(200, status)
}

func GetCurrentlyPlaying(c *gin.Context) {

}
