package controllers

import (
	"net/http"
	"net/url"
	"os"

	helper "github.com/fransedor/golang-spotify/helpers"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	client_id := os.Getenv("CLIENT_ID")
	redirect_uri := os.Getenv("REDIRECT_URI")

	state := helper.GenerateRandomString(16)
	scope := "user-read-private user-read-email user-read-playback-state"

	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", client_id)
	params.Add("redirect_uri", redirect_uri)
	params.Add("scope", scope)
	params.Add("state", state)
	location := url.URL{Scheme: "https", Host: "accounts.spotify.com", Path: "authorize", RawQuery: params.Encode()}
	c.Redirect(http.StatusFound, location.String())
}
