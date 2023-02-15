package controllers

import (
	"fmt"
	"log"
	"os"

	helper "github.com/fransedor/golang-spotify/helpers"
	"github.com/fransedor/golang-spotify/models"
	"github.com/gin-gonic/gin"
)

func GetAvailableDevices(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	spotifyURL := os.Getenv("SPOTIFY_URL")

	req, err := helper.CreateHTTPRequestWithHeader("GET", spotifyURL+"/me/player/devices", nil, authHeader[7:])
	if err != nil {
		log.Fatal(err)
	}

	var devices models.Devices
	_, err = helper.GetHTTPResponse(req, &devices)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available Devices: ", devices)
	c.IndentedJSON(200, devices)
}
