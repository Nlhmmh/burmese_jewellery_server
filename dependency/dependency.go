package dependency

import (
	"burmese_jewellery/env"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewDependency() error {
	if err := newDB(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.Get().DB.Host,
		env.Get().DB.Port,
		env.Get().DB.User,
		env.Get().DB.Password,
		env.Get().DB.Database,
	)); err != nil {
		return err
	}

	boil.SetDB(postgresDB)
	boil.DebugMode = true
	boil.DebugWriter = log.Logger

	return nil
}
