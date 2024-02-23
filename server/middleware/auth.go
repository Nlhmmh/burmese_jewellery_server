package middleware

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"

	"github.com/gin-gonic/gin"
)

// Auth - Auth JWT MiddleWare
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullPath := c.FullPath()

		if auth.CheckWhiteList(fullPath) || auth.CheckContainWhiteList(fullPath) {
			c.Next()
			return
		}

		token, err := auth.GetBearerToken(c)
		if err != nil {
			ers.UnAuthorized.New(err).Abort(c)
			return
		}

		claims, err := auth.ValidateToken(token)
		if err != nil {
			ers.UnAuthorized.New(err).Abort(c)
			return
		}
		claimsVal := *claims
		auth.SetUserID(c, claimsVal.UserID)

		if auth.CheckAdminList(fullPath) {
			if claims.Role == auth.RoleAdmin || claims.Role == auth.RoleStaff {
				c.Next()
				return
			}
			ers.UnAuthorizedAdminStaff.Abort(c)
			return
		}

		if auth.CheckAdminRoleAdminOnlyList(fullPath) {
			if claims.Role == auth.RoleAdmin {
				c.Next()
				return
			}
			ers.UnAuthorizedAdmin.Abort(c)
			return
		}

		if auth.CheckUserList(fullPath) {
			if claims.Role == auth.RoleUser {
				c.Next()
				return
			}
		}

		ers.UnAuthorized.Abort(c)
	}
}
