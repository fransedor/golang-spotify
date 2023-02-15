package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	client := &http.Client{}
	authHeader := c.Request.Header.Get("Authorization")

	spotifyURL := os.Getenv("SPOTIFY_URL")

	req, _ := http.NewRequest(http.MethodGet, spotifyURL+"/me/player/devices", nil)
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var devices APIReponse
	err = json.Unmarshal(bodyBytes, &devices)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available Devices: ", devices)
	c.IndentedJSON(200, devices)
}
