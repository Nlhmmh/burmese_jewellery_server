package dependency

import (
	"burmese_jewellery/config"
	"burmese_jewellery/env"
	"fmt"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/rs/zerolog/log"
)

// NewDependency - Initialize the database connection, set up the SQLBoiler configuration, and load the configuration.
// Singleton pattern is used to ensure that the database connection is only initialized once.
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

	config.Init()

	return nil
}
