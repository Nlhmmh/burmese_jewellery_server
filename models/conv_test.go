package models

import (
	"burmese_jewellery/orm"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestConvListFromORM(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ormList := []*orm.Gem{
			{
				GemID:     "6b6f9e68-37c2-4f25-8d48-2b0f8a7d03e1",
				Name:      "1",
				CreatedAt: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			},
		}

		gemID, err := uuid.Parse(ormList[0].GemID)
		assert.Nil(t, err)

		resp, err := ConvListFromORM(ormList, ConvGemFromORM)
		assert.Nil(t, err)
		assert.EqualValues(t, []*Gem{
			{
				GemId:     gemID,
				Name:      "1",
				CreatedAt: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			},
		}, resp)
	})

}
