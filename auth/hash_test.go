package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashPassword(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		hashPW, err := HashPassword("admin")
		assert.Nil(t, err)
		assert.NotEmpty(t, hashPW)
	})

}
