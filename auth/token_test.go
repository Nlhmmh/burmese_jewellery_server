package auth

import (
	"burmese_jewellery/config"
	"burmese_jewellery/orm"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateToken(t *testing.T) {

	token, err := GenerateToken("testUserID", orm.AccountAdminRoleStaff)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)

}

func Test_ValidateToken(t *testing.T) {

	token, err := GenerateToken("testUserID", orm.AccountAdminRoleStaff)
	assert.Nil(t, err)

	claims, err := ValidateToken(token)
	assert.Nil(t, err)
	assert.Equal(t, "testUserID", claims.UserID)
	assert.Equal(t, orm.AccountAdminRoleStaff.String(), claims.Role)
	assert.Equal(t, time.Now().Add(time.Hour*time.Duration(config.Get().JWTExpiredHours)).Unix(), claims.ExpiresAt)

}
