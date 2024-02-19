package models

import (
	"burmese_jewellery/orm"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

func ConvListFromORM[T any, V any](ormList []*T, convFromORMFunc func(*T) (*V, error)) ([]*V, error) {
	list := make([]*V, len(ormList))
	for i, v := range ormList {
		model, err := convFromORMFunc(v)
		if err != nil {
			return nil, err
		}
		list[i] = model
	}
	return list, nil
}

func ConvGemFromORM(orm *orm.Gem) (*Gem, error) {
	id, err := uuid.Parse(orm.GemID)
	if err != nil {
		return nil, err
	}
	g := &Gem{}
	g.GemId = id
	g.Name = orm.Name
	g.CreatedAt = orm.CreatedAt
	g.UpdatedAt = orm.UpdatedAt
	return g, nil
}

func ConvCategoryFromORM(orm *orm.Category) (*Category, error) {
	id, err := uuid.Parse(orm.CategoryID)
	if err != nil {
		return nil, err
	}
	c := &Category{}
	c.CategoryId = id
	c.Name = orm.Name
	c.Description = orm.Description
	c.ImageUrl = orm.ImageURL
	c.CreatedAt = orm.CreatedAt
	c.UpdatedAt = orm.UpdatedAt
	return c, nil
}

func ConvMaterialFromORM(orm *orm.Material) (*Material, error) {
	id, err := uuid.Parse(orm.MaterialID)
	if err != nil {
		return nil, err
	}
	m := &Material{}
	m.MaterialId = id
	m.Name = orm.Name
	m.CreatedAt = orm.CreatedAt
	m.UpdatedAt = orm.UpdatedAt
	return m, nil
}

func ConvJewelleryFromORM(orm *orm.Jewellery) (*Jewellery, error) {
	id, err := uuid.Parse(orm.JewelleryID)
	if err != nil {
		return nil, err
	}
	cID, err := uuid.Parse(orm.CategoryID)
	if err != nil {
		return nil, err
	}
	gID, err := uuid.Parse(orm.GemID)
	if err != nil {
		return nil, err
	}
	mID, err := uuid.Parse(orm.MaterialID)
	if err != nil {
		return nil, err
	}
	j := &Jewellery{}
	j.JewelleryId = id
	j.CategoryId = cID
	j.GemId = gID
	j.MaterialId = mID
	j.Name = orm.Name
	j.Description = orm.Description
	j.Price = orm.Price
	j.Quantity = orm.Quantity
	j.ImageUrl = orm.ImageURL
	j.IsPublished = orm.IsPublished
	j.CreatedAt = orm.CreatedAt
	j.UpdatedAt = orm.UpdatedAt
	return j, nil
}

func ConvAccountFromORM(orm *orm.Account) (*Account, error) {
	id, err := uuid.Parse(orm.AccountID)
	if err != nil {
		return nil, err
	}
	a := &Account{}
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
	a.AccountStatus = AccountStatus(orm.AccountStatus)
	a.CreatedAt = orm.CreatedAt
	a.UpdatedAt = orm.UpdatedAt
	return a, nil
}

func ConvAccountAdminFromORM(orm *orm.AccountAdmin) (*AccountAdmin, error) {
	id, err := uuid.Parse(orm.AccountAdminID)
	if err != nil {
		return nil, err
	}
	aa := &AccountAdmin{}
	aa.AccountAdminId = id
	aa.Mail = types.Email(orm.Mail)
	aa.AccountAdminRole = AccountAdminRole(orm.AccountAdminRole)
	aa.AccountAdminStatus = AccountAdminStatus(orm.AccountAdminStatus)
	aa.CreatedAt = orm.CreatedAt
	aa.UpdatedAt = orm.UpdatedAt
	return aa, nil
}

func ConvAccountProfileFromORM(orm *orm.AccountProfile) (*AccountProfile, error) {
	id, err := uuid.Parse(orm.AccountID)
	if err != nil {
		return nil, err
	}
	ap := &AccountProfile{}
	ap.AccountId = id
	ap.FirstName = orm.FirstName
	ap.LastName = orm.LastName
	if v := orm.Birthday; v.Valid {
		ap.Birthday = types.Date{Time: v.Time}
	}
	if v := orm.Gender; v.Valid {
		ap.Gender = Gender(v.Val)
	}
	ap.CreatedAt = orm.CreatedAt
	ap.UpdatedAt = orm.UpdatedAt
	return ap, nil
}

func ConvAccountWithProfileFromORM(ormm *orm.AccountWithProfile) (*AccountWithProfile, error) {
	awp := &AccountWithProfile{}

	a, err := ConvAccountFromORM(&orm.Account{
		AccountID:     ormm.AccountID,
		LoginType:     ormm.LoginType,
		LoginID:       ormm.LoginID,
		Mail:          ormm.Mail,
		Password:      ormm.Password,
		Phone:         ormm.Phone,
		AccountStatus: ormm.AccountStatus,
		CreatedAt:     ormm.CreatedAt,
		UpdatedAt:     ormm.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	awp.Account = *a

	ap, err := ConvAccountProfileFromORM(&orm.AccountProfile{
		AccountID: ormm.AccountID,
		FirstName: ormm.FirstName,
		LastName:  ormm.LastName,
		Birthday:  ormm.Birthday,
		Gender:    ormm.Gender,
		CreatedAt: ormm.CreatedAt,
		UpdatedAt: ormm.UpdatedAt,
	})
	if err != nil {
		return nil, err
	}
	awp.AccountProfile = *ap

	return awp, nil
}
