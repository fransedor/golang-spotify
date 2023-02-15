package controllers

import (
	"fmt"
	"log"
	"os"

	helper "github.com/fransedor/golang-spotify/helpers"
	"github.com/gin-gonic/gin"
)

type ExplicitContentStruct struct {
	Filter_enabled bool `json:"filter_enabled"`
	Filter_locked  bool `json:"filter_locked"`
}

type ExternalUrlsStruct struct {
	Spotify string `json:"spotify"`
}

type FollowersStruct struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type ImagesStruct struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Profile struct {
	Country          string `json:"country"`
	Display_name     string `json:"display_name"`
	Email            string `json:"email"`
	Explicit_content ExplicitContentStruct
	External_urls    ExternalUrlsStruct
	Followers        FollowersStruct
	Href             string `json:"href"`
	Id               string `json:"id"`
	Images           []ImagesStruct
	Product          string `json:"product"`
	Type             string `json:"type"`
	Uri              string `json:"uri"`
}

func GetProfile(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")

	spotifyURL := os.Getenv("SPOTIFY_URL")

	req, err := helper.CreateHTTPRequestWithHeader("GET", spotifyURL+"/me", nil, authHeader[7:])

	if err != nil {
		log.Fatal(err)
	}

	var profile Profile
	_, err = helper.GetHTTPResponse(req, &profile)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Profile: ", profile)
	c.IndentedJSON(200, profile)
}
