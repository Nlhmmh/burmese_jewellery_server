package models

import (
	"burmese_jewellery/orm"

	"github.com/google/uuid"
)

func (j *Jewellery) ConvFromORM(orm *orm.Jewellery) error {
	id, err := uuid.Parse(orm.JewelleryID)
	if err != nil {
		return err
	}
	j.JewelleryId = id
	j.CreatedAt = orm.CreatedAt
	j.UpdatedAt = orm.UpdatedAt
	return nil
}

func ConvJewelleryListFromORM(ormList []*orm.Jewellery) ([]*Jewellery, error) {
	list := make([]*Jewellery, len(ormList))
	for i, v := range ormList {
		list[i] = &Jewellery{}
		if err := list[i].ConvFromORM(v); err != nil {
			return nil, err
		}
	}
	return list, nil
}
