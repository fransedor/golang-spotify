package controllers

import (
	"fmt"
	"log"
	"os"

	helper "github.com/fransedor/golang-spotify/helpers"
	"github.com/fransedor/golang-spotify/models"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	spotifyURL := os.Getenv("SPOTIFY_URL")

	req, err := helper.CreateHTTPRequestWithHeader("GET", spotifyURL+"/me", nil, authHeader[7:])

	if err != nil {
		log.Fatal(err)
	}

	var profile models.Profile
	status, errorObj := helper.GetHTTPResponse(req, &profile)

	if status == "fail" {
		c.AbortWithStatusJSON(errorObj.Status, errorObj)
	}
	fmt.Println("Current Profile: ", profile)
	c.IndentedJSON(200, profile)
}
