package models

import (
	"burmese_jewellery/orm"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMapStructFields(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		source := &orm.AccountAdmin{
			AccountAdminID:     "3599ebb7-00a1-41c0-bdc8-8dcaff9f7b38",
			Mail:               "test@gmail.com",
			Password:           "",
			AccountAdminRole:   orm.AccountAdminRoleStaff,
			AccountAdminStatus: orm.AccountAdminStatusActive,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}
		target := &AccountAdmin{}
		val := MapStructFields(*source, *target).(AccountAdmin)
		*target = val

		assert.Equal(t, target.AccountAdminId.String(), source.AccountAdminID)
	})

}
