package auth

import (
	"burmese_jewellery/config"
	"burmese_jewellery/orm"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	jwtSecureKey = "secureSecretText"
)

type claims struct {
	UserID string
	Role   string
	jwt.StandardClaims
}

// GenerateToken - Generate Tokens
func GenerateToken(userID string, role orm.AccountAdminRole) (string, error) {

	// Create Token
	createdToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims{
			userID,
			role.String(),
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.Get().JWTExpiredHours)).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		},
	)

	// Sign Token
	token, err := createdToken.SignedString([]byte(jwtSecureKey))
	if err != nil {
		return "", err
	}

	return token, nil

}

// ValidateToken - Validate Token
func ValidateToken(encodedToken string) (*claims, error) {

	// Parse JWT
	token, err := jwt.ParseWithClaims(
		encodedToken,
		&claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecureKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	// IF token is invalid
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return token.Claims.(*claims), nil

}

func GetBearerToken(c *gin.Context) (string, error) {

	authHeader := c.GetHeader("Authorization")
	splittedTokenList := strings.Split(authHeader, " ")
	if len(splittedTokenList) < 2 {
		return "", errors.New("bearer token is wrong" + authHeader)
	}

	return splittedTokenList[1], nil

}
