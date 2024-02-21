package usecase

import (
	"burmese_jewellery/ers"
	"database/sql"
	"errors"
)

type FindFunc[T any] func() (T, error)
type InsertFunc func() error
type UpdateFunc[T any] func(T) error
type DelFunc func() error

func Upsert[T any](findFunc FindFunc[T], insertFunc InsertFunc, updateFunc UpdateFunc[T]) *ers.ErrResp {
	recordExists := true
	record, err := findFunc()
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return ers.InternalServer.New(err)
		}
		recordExists = false
	}

	if !recordExists {
		if err := insertFunc(); err != nil {
			return ers.InternalServer.New(err)
		}
	} else {
		if err := updateFunc(record); err != nil {
			return ers.InternalServer.New(err)
		}
	}

	return nil
}
