package handler

import (
	"burmese_jewellery/google_login"
	"burmese_jewellery/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	testGoogleLoginHTML = `
		<!DOCTYPE html>
		<html>
			<body>
				<a href="/api/auth/google/login">Login with Google</a>
			</body>
		</html>
	`
)

// TODO: For testing
// Display the main page
// (GET /api)
func (*Handler) GetApi(c *gin.Context) {
	if _, err := c.Writer.WriteString(testGoogleLoginHTML); err != nil {
		log.Error().Err(err)
		c.Redirect(http.StatusTemporaryRedirect, "/api")
	}
}

// Initiate Google OAuth2 login
// (GET /api/auth/google/login)
func (*Handler) GetApiAuthGoogleLogin(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, google_login.GetLoginURL())
}

// Handle Google OAuth2 callback
// (GET /api/auth/google/callback)
func (*Handler) GetApiAuthGoogleCallback(c *gin.Context, params models.GetApiAuthGoogleCallbackParams) {

	userInfo, err := google_login.Callback(c,
		params.State,
		params.Code,
	)
	if err != nil {
		log.Error().Err(err)
		c.Redirect(http.StatusTemporaryRedirect, "/api")
	}

	// TODO: Insert Accounts
	fmt.Println(userInfo)

	c.Redirect(http.StatusTemporaryRedirect, "/api")

}
