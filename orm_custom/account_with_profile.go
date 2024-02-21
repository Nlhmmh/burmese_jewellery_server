package orm_custom

import (
	"burmese_jewellery/orm"
	"time"

	"github.com/volatiletech/null/v8"
)

type AccountWithProfile struct {
	AccountID     string            `boil:"account_id" json:"accountID" toml:"accountID" yaml:"accountID"`
	LoginType     orm.LoginType     `boil:"login_type" json:"loginType" toml:"loginType" yaml:"loginType"`
	LoginID       null.String       `boil:"login_id" json:"loginID,omitempty" toml:"loginID" yaml:"loginID,omitempty"`
	Mail          null.String       `boil:"mail" json:"mail,omitempty" toml:"mail" yaml:"mail,omitempty"`
	Password      null.String       `boil:"password" json:"password,omitempty" toml:"password" yaml:"password,omitempty"`
	Phone         null.String       `boil:"phone" json:"phone,omitempty" toml:"phone" yaml:"phone,omitempty"`
	AccountStatus orm.AccountStatus `boil:"account_status" json:"accountStatus" toml:"accountStatus" yaml:"accountStatus"`
	CreatedAt     time.Time         `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt     time.Time         `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`

	FirstName string         `boil:"first_name" json:"firstName" toml:"firstName" yaml:"firstName"`
	LastName  string         `boil:"last_name" json:"lastName" toml:"lastName" yaml:"lastName"`
	Birthday  null.Time      `boil:"birthday" json:"birthday,omitempty" toml:"birthday" yaml:"birthday,omitempty"`
	Gender    orm.NullGender `boil:"gender" json:"gender,omitempty" toml:"gender" yaml:"gender,omitempty"`
}
