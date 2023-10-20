package middleware

import (
	"burmese_jewellery/auth"
	"burmese_jewellery/ers"
	"burmese_jewellery/orm"
	"errors"

	"github.com/gin-gonic/gin"
)

// Auth - Auth JWT MiddleWare
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullPath := c.FullPath()

		// Check White List
		if auth.CheckJWTWhiteList(fullPath) || auth.CheckContainWhiteList(fullPath) {
			c.Next()
			return
		}

		// Get token
		token, err := auth.GetBearerToken(c)
		if err != nil {
			ers.UnAuthorized.New(err).Abort(c)
			return
		}

		// Validate Token
		claims, err := auth.ValidateToken(token)
		if err != nil {
			ers.UnAuthorized.New(err).Abort(c)
			return
		}
		claimsVal := *claims
		c.Set("userID", claimsVal.UserID)

		// Check Admin
		if auth.CheckAdminList(fullPath) &&
			!(claims.Role == orm.AccountAdminRoleAdmin.String() ||
				claims.Role == orm.AccountAdminRoleStaff.String()) {
			ers.UnAuthorized.New(errors.New("not admin staff")).Abort(c)
			return
		}

		// Check Admin RoleAdmin
		if auth.CheckAdminRoleAdminOnlyList(fullPath) && claims.Role != orm.AccountAdminRoleAdmin.String() {
			ers.UnAuthorized.New(errors.New("not admin")).Abort(c)
			return
		}

		c.Next()
	}
}
