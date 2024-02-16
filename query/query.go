package query

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func Like(qList []qm.QueryMod, param *string, whereName string) []qm.QueryMod {
	if v := param; v != nil {
		qList = append(qList,
			qm.Where(
				fmt.Sprintf("%s LIKE ?", whereName),
				fmt.Sprintf("%%%v%%", lo.FromPtr(v)),
			),
		)
	}
	return qList
}

func EqUUID(qList []qm.QueryMod, param *uuid.UUID, whereName string) []qm.QueryMod {
	if v := param; v != nil {
		qList = append(qList,
			qm.Where(
				fmt.Sprintf("%s = ?", whereName),
				v.String(),
			),
		)
	}
	return qList
}

func Eq[T any](qList []qm.QueryMod, param *T, whereName string) []qm.QueryMod {
	if v := param; v != nil {
		qList = append(qList,
			qm.Where(
				fmt.Sprintf("%s = ?", whereName),
				lo.FromPtr(v),
			),
		)
	}
	return qList
}
