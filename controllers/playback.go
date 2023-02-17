package controllers

import (
	"log"
	"net/http"
	"os"

	helper "github.com/fransedor/golang-spotify/helpers"
	"github.com/fransedor/golang-spotify/models"
	"github.com/gin-gonic/gin"
)

func StartPlayback(c *gin.Context) {
	client := &http.Client{}
	authHeader := c.Request.Header.Get("Authorization")
	spotifyURL := os.Getenv("SPOTIFY_URL")
	req, err := helper.CreateHTTPRequestWithHeader("PUT", spotifyURL+"/me/player/play", nil, authHeader[7:])
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseObj := make(map[string]any)
	responseObj["status"] = 200
	responseObj["message"] = "started playback"

	c.IndentedJSON(200, responseObj)
}

func GetCurrentlyPlaying(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	spotifyURL := os.Getenv("SPOTIFY_URL")
	req, err := helper.CreateHTTPRequestWithHeader("GET", spotifyURL+"/me/player/currently-playing", nil, authHeader[7:])
	if err != nil {
		log.Fatal(err)
	}
	var currentTrack models.GetCurrentlyPlayingResponse
	_, err = helper.GetHTTPResponse(req, &currentTrack)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(200, currentTrack)
}
