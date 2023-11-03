package models

import (
	"burmese_jewellery/orm"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
)

func (aa *AccountAdmin) MapFromORM(orm *orm.AccountAdmin) error {
	if aaID, err := uuid.Parse(orm.AccountAdminID); err != nil {
		return err
	} else {
		aa.AccountAdminId = aaID
	}
	aa.Mail = types.Email(orm.Mail)
	aa.AccountAdminRole = AccountAdminAccountAdminRole(orm.AccountAdminRole)
	aa.AccountAdminStatus = AccountAdminAccountAdminStatus(orm.AccountAdminStatus)
	return nil
}

func MapAccountAdminListFromORM(ormList []*orm.AccountAdmin) ([]*AccountAdmin, error) {
	aaList := make([]*AccountAdmin, len(ormList))
	for i, v := range ormList {
		aaList[i] = &AccountAdmin{}
		if err := aaList[i].MapFromORM(v); err != nil {
			return nil, err
		}
	}
	return aaList, nil
}
