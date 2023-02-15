package controllers

import (
	"fmt"
	"log"
	"os"

	helper "github.com/fransedor/golang-spotify/helpers"
	"github.com/gin-gonic/gin"
)

type Device struct {
	Id                 string `json:"id"`
	Is_active          bool   `json:"is_active"`
	Is_private_session bool   `json:"is_private_session"`
	Is_restricted      bool   `json:"is_restricted"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	Volume_percent     int    `json:"volume_percent"`
}

type APIReponse struct {
	Devices []Device `json:"devices"`
}

func GetAvailableDevices(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	spotifyURL := os.Getenv("SPOTIFY_URL")

	req, err := helper.CreateHTTPRequestWithHeader("GET", spotifyURL+"/me/player/devices", nil, authHeader[7:])
	if err != nil {
		log.Fatal(err)
	}

	var devices APIReponse
	_, err = helper.GetHTTPResponse(req, &devices)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available Devices: ", devices)
	c.IndentedJSON(200, devices)
}
