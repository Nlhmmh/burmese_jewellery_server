package auth

import (
	"burmese_jewellery/orm"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateToken(t *testing.T) {

	token, err := GenerateToken("testUserID", Role(orm.AccountAdminRoleStaff))

	assert.Nil(t, err)
	assert.NotEmpty(t, token)

}

func Test_ValidateToken(t *testing.T) {
	tList := map[string]struct {
		userID string
		role   Role
		err    error
		setUp  func() string
	}{
		"ok/admin": {
			userID: "testUserIDAdmin",
			role:   RoleAdmin,
			err:    nil,
			setUp: func() string {
				token, err := GenerateToken("testUserIDAdmin", RoleAdmin)
				assert.Nil(t, err)
				return token
			},
		},
		"ok/staff": {
			userID: "testUserIDStaff",
			role:   RoleStaff,
			err:    nil,
			setUp: func() string {
				token, err := GenerateToken("testUserIDStaff", RoleStaff)
				assert.Nil(t, err)
				return token
			},
		},
		"error": {
			userID: "",
			role:   RoleAdmin,
			err:    errors.New("token contains an invalid number of segments"),
			setUp: func() string {
				return "invalidToken"
			},
		},
	}

	for name, tt := range tList {
		t.Run(name, func(t *testing.T) {
			token := tt.setUp()
			claims, err := ValidateToken(token)
			if err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
			}
			if claims != nil {
				assert.Equal(t, tt.userID, claims.UserID)
				assert.Equal(t, tt.role, claims.Role)
			}
		})
	}
}
