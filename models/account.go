package models

import (
	"burmese_jewellery/orm"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

func (a *Account) ConvFromORM(orm *orm.Account) error {
	id, err := uuid.Parse(orm.AccountID)
	if err != nil {
		return err
	}
	a.AccountId = id
	a.LoginType = AccountLoginType(orm.LoginType)
	if v := orm.LoginID; v.Valid {
		vv := v.String
		a.LoginId = &vv
	}
	if v := orm.Mail; v.Valid {
		vv := types.Email(v.String)
		a.Mail = &vv
	}
	a.AccountStatus = AccountAccountStatus(orm.AccountStatus)
	a.CreatedAt = orm.CreatedAt
	a.UpdatedAt = orm.UpdatedAt
	return nil
}

func ConvAccountListFromORM(ormList []*orm.Account) ([]*Account, error) {
	list := make([]*Account, len(ormList))
	for i, v := range ormList {
		list[i] = &Account{}
		if err := list[i].ConvFromORM(v); err != nil {
			return nil, err
		}
	}
	return list, nil
}
