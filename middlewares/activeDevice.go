package middlewares

import (
	"fmt"
	"log"
	"os"

	helper "github.com/fransedor/golang-spotify/helpers"
	"github.com/fransedor/golang-spotify/models"
	"github.com/gin-gonic/gin"
)

func GetActiveDevice(c *gin.Context) {
	fmt.Println("Running GetActiveDevice")
	authHeader := c.Request.Header.Get("Authorization")

	spotifyURL := os.Getenv("SPOTIFY_URL")

	req, err := helper.CreateHTTPRequestWithHeader("GET", spotifyURL+"/me/player/devices", nil, authHeader[7:])
	if err != nil {
		log.Fatal(err)
	}

	var devices models.Devices
	status, errorObj := helper.GetHTTPResponse(req, &devices)
	if status == "fail" {
		c.AbortWithStatusJSON(errorObj.Status, errorObj)
		return
	}

	var hasActiveDevice bool = false
	if len(devices.Devices) > 0 {
		hasActiveDevice = true
	}
	fmt.Println("Checking active device: ", hasActiveDevice)
	if !hasActiveDevice {
		c.AbortWithStatusJSON(400, helper.CreateErrorResponse(400, "Open your spotify player"))
	} else {
		c.Next()
	}
}
