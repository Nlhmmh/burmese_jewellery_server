package models

import (
	"burmese_jewellery/orm"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

func (aa *AccountAdmin) ConvFromORM(orm *orm.AccountAdmin) error {
	aaID, err := uuid.Parse(orm.AccountAdminID)
	if err != nil {
		return err
	}
	aa.AccountAdminId = aaID
	aa.Mail = types.Email(orm.Mail)
	aa.AccountAdminRole = AccountAdminAccountAdminRole(orm.AccountAdminRole)
	aa.AccountAdminStatus = AccountAdminAccountAdminStatus(orm.AccountAdminStatus)
	aa.CreatedAt = orm.CreatedAt
	aa.UpdatedAt = orm.UpdatedAt
	return nil
}

func ConvAccountAdminListFromORM(ormList []*orm.AccountAdmin) ([]*AccountAdmin, error) {
	aaList := make([]*AccountAdmin, len(ormList))
	for i, v := range ormList {
		aaList[i] = &AccountAdmin{}
		if err := aaList[i].ConvFromORM(v); err != nil {
			return nil, err
		}
	}
	return aaList, nil
}
