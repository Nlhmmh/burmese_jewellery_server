package handler

import (
	"burmese_jewellery/ers"
	"burmese_jewellery/google_login"
	"burmese_jewellery/models"
	"burmese_jewellery/orm"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
		ers.TmpRedirect.New(err).TmpRedirect(c, "/api")
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
		ers.TmpRedirect.New(err).TmpRedirect(c, "/api")
		return
	}

	a, err := orm.Accounts(
		orm.AccountWhere.Mail.EQ(null.StringFrom(userInfo.Email)),
	).OneG(c)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		ers.InternalServer.New(err).Abort(c)
		return
	}
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		a = &orm.Account{
			AccountID:     uuid.New().String(),
			LoginType:     orm.LoginTypeGoogle,
			LoginID:       null.StringFrom(userInfo.ID),
			Mail:          null.StringFrom(userInfo.Email),
			AccountStatus: orm.AccountStatusActive,
		}
		if err := a.InsertG(c, boil.Infer()); err != nil {
			ers.InternalServer.New(err).Abort(c)
			return
		}
	}

	apExists, err := orm.AccountProfileExistsG(c, a.AccountID)
	if err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	resp := &models.Account{}
	if err := resp.ConvFromORM(a); err != nil {
		ers.InternalServer.New(err).Abort(c)
		return
	}

	c.JSON(http.StatusOK, &models.AuthGoogleCallbackResp{
		Account:      resp,
		IsRegistered: apExists,
	})
}
