package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTracks(access_token *string, refresh_token *string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("access token: " + *access_token)
		fmt.Println("refresh token: " + *refresh_token)
	}
}
