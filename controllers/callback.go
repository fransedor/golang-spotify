package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type GetAccessTokenRequestBody struct {
	Code         string
	Redirect_uri string
	Grant_type   string
}

type GetAccessTokenResponseBody struct {
	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Scope         string `json:"scope"`
	Expires_in    int    `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
}

func Callback(access_token *string, refresh_token *string) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := &http.Client{}

		client_id := os.Getenv("CLIENT_ID")
		client_secret := os.Getenv("CLIENT_SECRET")
		redirectURI := os.Getenv("REDIRECT_URI")

		requestURL := c.Request.URL.Query()
		code := requestURL.Get("code")

		authHeaderStr := client_id + ":" + client_secret
		encodedAuth := base64.StdEncoding.EncodeToString([]byte(authHeaderStr))

		spotifyTokenRequestBody := url.Values{}
		spotifyTokenRequestBody.Set("code", code)
		spotifyTokenRequestBody.Set("redirect_uri", redirectURI)
		spotifyTokenRequestBody.Set("grant_type", "authorization_code")
		encodedBody := spotifyTokenRequestBody.Encode()

		req, _ := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(encodedBody))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", "Basic "+encodedAuth)
		response, err := client.Do(req)
		if err != nil {
			log.Fatalf("error get access token")
		}
		defer response.Body.Close()

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var accessTokenResponse GetAccessTokenResponseBody
		err = json.Unmarshal(bodyBytes, &accessTokenResponse)
		if err != nil {
			log.Fatal(err)
		}
		*access_token = accessTokenResponse.Access_token
		*refresh_token = accessTokenResponse.Refresh_token
	}

}
